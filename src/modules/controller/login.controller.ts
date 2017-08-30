import { UsersService } from "./../service/users.service"
import { Controller, Get, Post } from '@nestjs/common'
import { NextFunction } from 'express'
import { Response, Request, Param, Body } from '@nestjs/common'

@Controller('login')
export class LoginController {
  constructor(private UsersService: UsersService) {

  }
  @Post()
  async login( @Response() res, @Body() user) {
    console.log('req.body:', user)
    let code = 0
    this.UsersService.verifyUser(user)
      .then(
      _user => res.send({ code: 1 }),
      err_code => res.send({ code: err_code })
      )

  }
}