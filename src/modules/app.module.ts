import { UsersService } from './service/users.service'
import { UsersController } from './controller/users.controller'
import { LoginController } from './controller/login.controller'
import { Module } from '@nestjs/common'

@Module({
    modules: [],
    controllers: [UsersController, LoginController],
    components: [UsersService]
})
export class ApplicationModule { }