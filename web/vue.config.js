const { defineConfig } = require("@vue/cli-service");
const CompressionWebpackPlugin = require('compression-webpack-plugin');
const isProdOrTest = process.env.NODE_ENV !== 'development'
module.exports = defineConfig({
    productionSourceMap: false,// 设为false，既可以减少包大小，也可以加密源码
    lintOnSave: false,
    transpileDependencies: true,
    publicPath: "./",
    outputDir: "dist", // 打包到根目录的dist文件夹下
    chainWebpack: (config) => {
        config.plugins.delete('prefetch');//默认开启prefetch(预先加载模块)，提前获取用户未来可能会访问的内容 在首屏会把这十几个路由文件，都一口气下载了 所以我们要关闭这个功能模块
        config.plugin("html").tap((args) => {
            args[0].title = "项目管理";
            return args;
        });
        if (isProdOrTest) {
            // 对超过10kb的文件gzip压缩
            config.plugin('compressionPlugin').use(
                new CompressionWebpackPlugin({
                    test: /\.(js|css|html)$/,// 匹配文件名
                    threshold: 10240,
                })
            );
        }
    },
    devServer: {
        hot: true, //是否自动保存
        open: true, //是否自动启动
        port: 5001,
        historyApiFallback: true,
        https: false, //是否为https 请求 https:{type:Boolean}
        proxy: {
            "/mark": {
                target: "http://localhost:8080",
                changOrigin: true, //允许跨域
                ws: false,
                pathRewrite: {
                    "^/mark":
                        "" /* 重写路径，当我们在浏览器中看到请求的地址为：http://localhost:8080/api/core/getData/userInfo 时
                    实际上访问的地址是：http://121.121.67.254:8185/core/getData/userInfo,因为重写了 /api
                   */
                }
            },
            "/": {
                target: "http://localhost:8080",
                changOrigin: true, //允许跨域
                ws: false,
                pathRewrite: {
                    "^/": "" /* 重写路径，当我们在浏览器中看到请求的地址为：http://localhost:8080/api/core/getData/userInfo 时
                    实际上访问的地址是：http://121.121.67.254:8185/core/getData/userInfo,因为重写了 /api
                   */
                }
            }
        }
    }
});
