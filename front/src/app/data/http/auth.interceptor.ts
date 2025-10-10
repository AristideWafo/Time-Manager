import { HttpHandlerFn, HttpRequest } from '@angular/common/http';

export function authTokenInterceptor(req: HttpRequest<unknown>, next: HttpHandlerFn) {
    // Do not add Authorization for login endpoint
    if (req.url.endsWith('/authentification')) {
        return next(req);
    }

    const storedToken = localStorage.getItem('access_token');
    const reqWithAuth = storedToken
        ? req.clone({ setHeaders: { Authorization: `Bearer ${storedToken}` } })
        : req;

    return next(reqWithAuth);
}


