function createFakeUserList() {
  return [
    {
      userId: '1',
      userName: 'jyiot',
      nickName: '美男子',
      avatar: 'https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640',
      password: '123456',
      token: 'fakeToken1',
      homePath: '/dashboard',
      roles: [
        {
          name: 'Super Admin',
          value: 'super',
        },
      ],
      email: 'test@gmail.com',
      address: 'Xiamen City 77',
      phone: '0592-268888888',
    },
    {
      userId: '2',
      userName: 'test',
      password: '123456',
      nickName: 'test user',
      avatar: 'https://q1.qlogo.cn/g?b=qq&nk=339449197&s=640',
      token: 'fakeToken2',
      homePath: '/dashboard/workbench',
      roles: [
        {
          name: 'Tester',
          value: 'test',
        },
      ],
      email: 'test@gmail.com',
      address: 'Xiamen City 77',
      phone: '0592-268888888',
    },
  ];
}

export default class UserService {
  async Login() {
    return createFakeUserList();
  }

  async Logout() {}

  async GetMe() {}
}
