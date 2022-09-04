import Koa from 'koa';
import KoaRouter from 'koa-router';
import KoaCors from 'koa2-cors';
import KoaStatic from 'koa-static';
import path from 'path';
import { env } from 'node:process';
import dotenv from 'dotenv';

import { AppRoutes } from './config/routes';

// 加载环境变量
dotenv.config({ path: '.env.' + env.NODE_ENV });

const app = new Koa();
const router = new KoaRouter();

app.use(KoaCors());
app.use(router.routes());
app.use(router.allowedMethods());
app.use(KoaStatic(path.join(__dirname)));

// 加载路由
AppRoutes.forEach((route) =>
  router[route.method](env.MOCK_SERVER_BASE_URI + route.path, route.action),
);

// 监听端口
app.listen(env.PORT, () => {
  console.log(`server listening on: ${env.MOCK_SERVER_PORT}`);
});

module.exports = app;
