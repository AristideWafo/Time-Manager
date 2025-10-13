import { HttpHandlerFn, HttpRequest } from '@angular/common/http';

export function authTokenInterceptor(req: HttpRequest<unknown>, next: HttpHandlerFn) {
    if (req.url.endsWith('/authentification')) return next(req);

    const token = localStorage.getItem('access_token');
    const reqWithAuth = token ? req.clone({ setHeaders: { Authorization: `Bearer ${token}` } }) : req;
    return next(reqWithAuth);
}
