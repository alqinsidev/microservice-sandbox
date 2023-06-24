import { JwtService } from '@nestjs/jwt';
export declare class JwtAuthService {
    private readonly jwtService;
    constructor(jwtService: JwtService);
    verifyToken(token: string): Promise<any>;
}
