import 'dotenv/config';
import { JwtAuthService } from '../../modules/jwt/jwt.service';
declare const JwtStrategy_base: new (...args: any[]) => any;
export declare class JwtStrategy extends JwtStrategy_base {
    private readonly jwtAuthService;
    constructor(jwtAuthService: JwtAuthService);
    validate(payload: any): Promise<any>;
}
export {};
