import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';

export function roleGuard(requiredRole: string): CanActivateFn {
    return () => {
        const router = inject(Router);
        const role = localStorage.getItem('Role');

        if (role === requiredRole) {
            return true;
        }

        router.navigate(['/login']);
        return false;
    };
}
