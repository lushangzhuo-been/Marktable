# 阶段1：构建
# 使用更稳定的镜像并安装构建工具
FROM node:18-bullseye-slim AS builder
WORKDIR /app
# 调试：检查 node 和 npm 版本
RUN node -v && npm -v

# 接收构建参数
ARG NODE_ENV=out

# 将 ARG 转换为 ENV，供 npm run build 使用
ENV NODE_ENV=${NODE_ENV}

COPY package*.json ./
# 安装依赖（严格锁定版本）
RUN npm ci
COPY . .
RUN npm run build           # 执行构建命令

# 阶段2：运行
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
# 如果 nginx.conf 在项目根目录
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80

# 正确配置 保持 Nginx 前台运行
CMD ["nginx", "-g", "daemon off;"]