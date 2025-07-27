# 项目名称
MarkTable

# 项目简介
MarkTable 是一款高度灵活的项目管理平台，通过多维自定义配置（如流程、权限、视图、报表），
支持企业快速搭建适配不同业务场景的数字化工作流，涵盖研发管理、客户成功、项目协作等领域，
助力团队实现跨部门高效协同与流程标准化

# 平台介绍
http://47.109.66.52/#/login


# MarkTable Docker化部署
    
    1. 概述
    本文档描述如何使用 Docker Compose 部署 MarkTable 的完整环境。
       所有服务在容器化，通过 docker-compose.yml 一键部署。
    2. 前提条件
        - 已安装 Docker（版本 ≥ 20.10.0）
        - 已安装 Docker Compose（版本 ≥ 2.0.0）
    - 服务器资源：
    - CPU: 至少 4 核
      - 内存: 至少 8GB
      - 磁盘: 根据数据需求预留空间

    3. 部署步骤
        3.0 前端代码地址 https://github.com/lushangzhuo-been/mark3_client
        3.1 将docker-compose.yml文件copy出项目
        3.2 启动服务 docker-compose up -d
        3.3 验证部署 docker-compose ps
    
        
        
        
        