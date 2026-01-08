import { TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { authGuard } from '../../view/guards/auth.guard';
import { ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

describe('authGuard', () => {
    let router: jasmine.SpyObj<Router>;
    const dummyRoute = {} as ActivatedRouteSnapshot;
    const dummyState = {} as RouterStateSnapshot;

    beforeEach(() => {
        router = jasmine.createSpyObj('Router', ['navigate']);
        TestBed.configureTestingModule({
            providers: [
                { provide: Router, useValue: router }
            ]
        });
    });

    it('should allow access if token exists', () => {
        spyOn(localStorage, 'getItem').and.returnValue('fake-token');
        const result = TestBed.runInInjectionContext(() =>
            authGuard(dummyRoute, dummyState)
        );
        expect(result).toBeTrue();
        expect(router.navigate).not.toHaveBeenCalled();
    });

    it('should redirect to /login if token is missing', () => {
        spyOn(localStorage, 'getItem').and.returnValue(null);
        const result = TestBed.runInInjectionContext(() =>
            authGuard(dummyRoute, dummyState)
        );
        expect(result).toBeFalse();
        expect(router.navigate).toHaveBeenCalledWith(['/login']);
    });
});
