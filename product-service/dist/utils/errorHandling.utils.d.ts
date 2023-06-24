export declare class NotFoundError extends Error {
    constructor(message: string);
}
export declare class BadRequestError extends Error {
    constructor(message: string);
}
export declare class ForbiddenError extends Error {
    constructor(message: string);
}
declare const HandleErrorException: (error: any) => never;
export default HandleErrorException;
