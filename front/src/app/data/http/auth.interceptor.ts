import { HttpHandlerFn, HttpRequest } from '@angular/common/http';

export function authTokenInterceptor(req: HttpRequest<unknown>, next: HttpHandlerFn) {
    console.log("➡️ Interceptor : URL =", req.url);

    if (req.url.endsWith('/authentification')) {
        return next(req);
    }

    const token = localStorage.getItem('access_token');
    let clonedReq = req;

    if (token) {
        clonedReq = req.clone({
            headers: req.headers.set('api_token', token)
        });
    }

    return next(clonedReq);
}
