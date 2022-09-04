import { Result } from '../utils/utils';
import * as protocol from '../../../../api/user';

function createFakeUserList() {
  return [
    {
      userId: 100,
      userName: 'kratos',
      nickName: '美男子',
      avatar: 'https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640',
      email: 'test@gmail.com',
      address: 'Xiamen City 77',
      phone: '0592-268888888',
    },
    {
      userId: 200,
      userName: 'test',
      nickName: 'test user',
      avatar: 'https://q1.qlogo.cn/g?b=qq&nk=339449197&s=640',
      email: 'test@gmail.com',
      address: 'Xiamen City 77',
      phone: '0592-268888888',
    },
  ];
}

export default class UserService {
  async Login(ctx) {
    const fakeUserInfo = createFakeUserList()[0];
    const params: protocol.LoginResponse = {
      id: fakeUserInfo.userId,
      userName: fakeUserInfo.userName,
      token: 'dfsdfsdfsdfsdfdf',
    };
    ctx.body = Result.success(params);
  }

  async Logout(ctx) {
    ctx.body = Result.success({});
  }

  async GetMe(ctx) {
    const fakeUserInfo = createFakeUserList()[0];
    ctx.body = Result.success(fakeUserInfo);
  }
}
