import Koa from 'koa';
import path from 'path';
import Router from 'koa-router';
import cors from 'koa2-cors';
import koaStatic from 'koa-static';

import { AppRoutes } from './config/routes';
import { PORT, baseUri } from './config/config';

//console.log(AppRoutes)

const app = new Koa();

const router = new Router();

app.use(cors());
app.use(router.routes());
app.use(router.allowedMethods());
app.use(koaStatic(path.join(__dirname)));

// 加载路由
AppRoutes.forEach((route) => router[route.method](baseUri+route.path, route.action));

app.listen(PORT, () => {
  console.log(`server listening at: ${PORT}`);
});
