import { HttpHandlerFn, HttpRequest } from '@angular/common/http';

export function authTokenInterceptor(req: HttpRequest<unknown>, next: HttpHandlerFn) {
    console.log("➡️ Interceptor : URL =", req.url);

    if (req.url.endsWith('/authentification')) return next(req);

    const token = localStorage.getItem('access_token');
    if (token) {
        console.log("🔑 Interceptor : ajout du token", token);
        const reqWithAuth = req.clone({ setHeaders: { Authorization: `Bearer ${token}` } });
        return next(reqWithAuth);
    } else {
        console.warn("⚠️ Interceptor : pas de token trouvé");
    }

    return next(req);
}

