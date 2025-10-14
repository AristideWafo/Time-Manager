import { HttpHandlerFn, HttpRequest } from '@angular/common/http';

export function authTokenInterceptor(req: HttpRequest<unknown>, next: HttpHandlerFn) {
    console.log("➡️ Interceptor : URL =", req.url);

    // Ignore login request
    if (req.url.endsWith('/authentification')) {
        return next(req);
    }

    const token = localStorage.getItem('access_token');
    let clonedReq = req;

    if (token) {
        clonedReq = req.clone({
            headers: req.headers.set('api_token', token)
        });
        console.log("🔑 Interceptor : ajout du token dans api_token", token);
    } else {
        console.warn("⚠️ Interceptor : pas de token trouvé");
    }

    return next(clonedReq);
}
