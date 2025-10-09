import { HttpHandlerFn, HttpRequest } from '@angular/common/http';

export function authTokenInterceptor(req: HttpRequest<unknown>, next: HttpHandlerFn) {
    const storedToken = localStorage.getItem('access_token');
    const tokenToUse = storedToken ?? 'FAKE_TOKEN_DEV_123';

    const reqWithAuth = tokenToUse
        ? req.clone({ setHeaders: { Authorization: `Bearer ${tokenToUse}` } })
        : req;

    console.log(tokenToUse);

    return next(reqWithAuth);
}


