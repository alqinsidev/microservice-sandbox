"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.JwtAuthModule = void 0;
const common_1 = require("@nestjs/common");
const jwt_1 = require("@nestjs/jwt");
const jwt_service_1 = require("./jwt.service");
const jwt_strategy_1 = require("../../guards/Jwt/jwt.strategy");
const passport_1 = require("@nestjs/passport");
const config_1 = require("@nestjs/config");
let JwtAuthModule = class JwtAuthModule {
};
JwtAuthModule = __decorate([
    (0, common_1.Module)({
        imports: [
            jwt_1.JwtModule.registerAsync({
                useFactory: async (configService) => ({
                    secret: configService.get('JWT_ACCESS_SECRET_KEY'),
                }),
                inject: [config_1.ConfigService],
            }),
            passport_1.PassportModule,
        ],
        providers: [jwt_service_1.JwtAuthService, jwt_strategy_1.JwtStrategy],
        exports: [jwt_service_1.JwtAuthService],
    })
], JwtAuthModule);
exports.JwtAuthModule = JwtAuthModule;
//# sourceMappingURL=jwt.module.js.map