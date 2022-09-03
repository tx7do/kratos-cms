const { name } = require('./package.json');
const path = require('path');

module.exports = {
  apps: [
    {
      // 应用程序名称
      name,
      // 执行文件
      script: path.resolve(__dirname, './dist/index.js'),
      // 应用启动实例个数，仅在cluster模式有效 默认为fork；或者 max
      instances: require('os').cpus().length,
      // 启用/禁用应用程序崩溃或退出时自动重启，默认为true, 发生异常的情况下自动重启
      autorestart: true,
      // 是否启用监控模式，默认是false。如果设置成true，当应用程序变动时，pm2会自动重载。这里也可以设置你要监控的文件。
      watch: true,
      env_production: {
        NODE_ENV: 'production',
        PORT: 8800,
      },
    },
  ],
};
