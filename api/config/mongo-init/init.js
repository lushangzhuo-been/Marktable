// 创建应用用户
db.createUser({
    user: "root",
    pwd: "example",
    roles: [
        { role: "readWrite", db: "mark3" },
        { role: "read", db: "reporting" }
    ]
});

// 创建初始集合和数据
db = db.getSiblingDB('mark3');

// 创建messages集合
if (!db.getCollectionNames().includes("messages")) {
    db.createCollection("messages");

    // 插入初始数据
    db.messages.insertMany([
        {
            title: "Welcome",
            content: "Initial setup complete",
            created_at: new Date()
        },
        {
            title: "Reminder",
            content: "Configure your application",
            created_at: new Date()
        }
    ]);
}

// 创建索引
db.messages.createIndex({ title: 1 });