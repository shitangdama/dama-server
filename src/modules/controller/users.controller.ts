import { UsersService } from "./../service/users.service"
import { Controller, Get, Post } from '@nestjs/common';
import { Request, NextFunction } from 'express'
import { Response, Param, Body } from '@nestjs/common';

@Controller('users')
export class UsersController {
  constructor(private usersService: UsersService) { }

  @Get()
  async getAllUsers( @Response() res) {
    const users = await this.usersService.getAllUsers();
    res.status(200).json(users);
  }

  @Get('/:id')
  async getUser( @Response() res, @Param('id') id) {
    const user = await this.usersService.getUserById(+id);
    res.status(200).json(user);
  }

  @Post()
  async addUser( @Response() res, @Body('user') user) {
    const msg = await this.usersService.getUserById(user);
    res.status(200).json(msg);
  }
}