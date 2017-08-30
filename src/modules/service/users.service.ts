import { Component } from '@nestjs/common';
import { HttpException } from '@nestjs/core';

@Component()
export class UsersService {
  private users = [
    { id: 1, name: "dama", password: "aichirou" },
    { id: 2, name: "jader", password: "havefun" },
    { id: 3, name: "admin", password: "admin" },
  ];
  getAllUsers() {
    return Promise.resolve(this.users);
  }
  getUserById(id: number) {
    const user = this.users.find((user) => user.id === id);
    if (!user) {
      throw new HttpException("User not found", 404);
    }
    return Promise.resolve(user);
  }
  verifyUser(user: any): Promise<any> {
    const _user = this.users.find(user_ => user_.name === user.name);
    let ret
    if (!_user) {
      ret = Promise.reject(0)
    } else if (_user.password !== user.password) {
      ret = Promise.reject(0)
    } else {
      ret = Promise.resolve(user);
    }
    return ret

  }
  addUser(user) {
    this.users.push(user);
    return Promise.resolve();
  }
}