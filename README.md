![cover-v5-optimized](./images/app.png)


# 平台介绍
marktable 是一款高度灵活的项目管理平台，通过多维自定义配置（如流程、权限、视图、报表），
支持企业快速搭建适配不同业务场景的数字化工作流，涵盖研发管理、客户成功、项目协作等领域，
助力团队实现跨部门高效协同与流程标准化。以下是其核心功能列表：
</br> </br>
**1. 自定义流程配置**:
平台支持深度自定义流程配置，满足您的独特需求。

**2. 自定义视图**:
聚焦关键指标，告别信息过载。

**3. 自动化规则**:
告别繁琐重复的手动操作，让 marktable 的智能规则引擎为您代劳

**3. 强大的数据可视化**:
告别枯燥数字的困扰，marktable 强大的可视化功能助您轻松解读数据

**4. 多种格式文档在线预览**:
终结文档分散、查找困难的局面，marktable 提供强大的文档集成与协作能力

# 功能比较
<table data-draft-node="block" data-draft-type="table" data-size="normal" data-row-style="normal">
    <tbody>
    <tr>
        <th>功能维度</th>
        <th>marktable</th>
        <th>Jira</th>
        <th>Teambition</th>
        <th>Trello</th>
    </tr>
    <tr>
        <td>自定义流程</td>
        <td>✅ 全流程可视化搭建（支持多工作流并行）</td>
        <td>⚠需代码配置复杂逻辑</td>
        <td>✅ 基础流程自定义</td>
        <td>❌ 仅看板列层级调整</td>
    </tr>
    <tr>
        <td>自定义视图</td>
        <td>✅ 多维视图</td>
        <td>❌ 视图固化</td>
        <td>✅ 基础视图筛选</td>
        <td>❌ 仅看板视图</td>
    </tr>
    <tr>
        <td>自动化规则</td>
        <td>✅ 无代码灵活触发（状态/字段/日期联动）</td>
        <td>✅ 强大但需脚本基础</td>
        <td>⚠ 基础规则（有限触发）</td>
        <td>❌ 仅付费版支持简单规则</td>
    </tr>
    <tr>
        <td>数据可视化</td>
        <td>✅ 多种图表类型（一键生成）</td>
        <td>✅ 需安装插件</td>
        <td>✅ 基础报表</td>
        <td>❌ 无原生图表</td>
    </tr>
    <tr>
        <td>文档协作整合</td>
        <td>✅ 多格式在线预览（PDF/Word/Excel等）</td>
        <td>❌ 仅链接跳转</td>
        <td>✅ 关联钉钉文档</td>
        <td>❌ 仅基础附件功能</td>
    </tr>
    <tr>
        <td>上手门槛</td>
        <td>⭐⭐（配置即用）</td>
        <td>⭐⭐⭐⭐（需专业培训）</td>
        <td>⭐⭐（界面友好）</td>
        <td>⭐（极简）</td>
    </tr>
    <tr>
        <td>成本</td>
        <td>✅ 完全免费，无隐藏费用</td>
        <td>⚠ 最多支持10个用户（缺报表/自动化）插件费用高（图表插件10＄/人）</td>
        <td>✅ 20人以下可用基础任务/文档存储扩容贵（100GB以上50/月）</td>
        <td>✅ 基础看板文件预览弱（无Office在线预览）</td>
    </tr>
    </tbody>
</table>

# 使用 marktable

- **云 </br>**
  我们提供[ marktable 试用版](https://marktable.cn/)，任何人都可以尝试登陆体验 marktable 提供的功能。

- **自托管 marktable 社区版</br>**
  使用 Docker 部署，快速在您的环境中运行 marktable。
  使用我们的[文档](http://marktable.cn:8084/src/md/%E4%BA%A7%E5%93%81%E6%A6%82%E8%BF%B0.html)进行进一步的参考和更深入的说明。

- **面向企业/组织的 marktable</br>**
  我们提供额外的面向企业的功能。[给我们发送电子邮件](mailto:360826018@qq.com)讨论企业需求。 </br>

# 安装社区版
**1. 概述**:
- 使用 Docker Compose 部署 marktable 的完整环境。所有服务在容器化，通过 docker-compose.yml 一键部署。

**2. 前提条件**:
- 已安装 Docker（版本 ≥ 20.10.0）
- 已安装 Docker Compose（版本 ≥ 2.0.0）
- 服务器资源：
    - CPU: 至少 4 核
    - 内存: 至少 8GB
    - 磁盘: 根据数据需求预留空间

**3. 部署步骤**:
- 3.1 执行命令构建服务
  ```
  docker-compose build
  ```
- 3.2 启动服务
  ```
    docker-compose up -d
  ```

- 3.3 验证部署
  ```
  docker-compose ps
  ```
# 社区与支持
我们欢迎您为 marktable 做出贡献，以帮助改善 marktable。包括：提交代码、问题、新想法，或分享您基于 marktable 创建的有趣且有用的 AI 应用程序。同时，我们也欢迎您在不同的活动、会议和社交媒体上分享 marktable。

- [GitHub Discussion](https://github.com/marktable-ai/marktable/discussions/)👉：分享您的应用程序并与社区交流。
- [GitHub Issues](https://github.com/marktable-ai/marktable/issues)👉：使用 marktable 时遇到的错误和问题，请参阅[贡献指南](CONTRIBUTING_CN.md)。
- [电子邮件支持](mailto:360826018@qq.com)👉：关于使用 marktable 的问题。
- (QQ群): 1031597375👉：分享您的应用程序并与社区交流。
- [商业许可](mailto:360826018@qq.com)👉：有关商业用途许可 marktable 的商业咨询。

# License
本仓库遵循 [marktable Open Source License](LICENSE) 开源协议，该许可证本质上是 GNU AFFERO GENERAL PUBLIC LICENSE Version 3。
主要义务:
- 分发软件时：必须提供源代码和许可证文本。

- 作为网络服务运行时：需向用户提供下载对应修改版源代码的途径。

- 修改代码后：衍生作品必须保持AGPLv3授权。
