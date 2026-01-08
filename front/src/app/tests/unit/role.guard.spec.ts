import { TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { roleGuard } from '../../view/guards/role.guard';
import { ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

describe('roleGuard', () => {
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

    it('should allow access if role matches', () => {
        spyOn(localStorage, 'getItem').and.returnValue('ADMIN');

        const guard = roleGuard('ADMIN');
        const result = TestBed.runInInjectionContext(() =>
            guard(dummyRoute, dummyState)
        );

        expect(result).toBeTrue();
        expect(router.navigate).not.toHaveBeenCalled();
    });

    it('should redirect to /login if role does not match', () => {
        spyOn(localStorage, 'getItem').and.returnValue('EMPLOYEE');

        const guard = roleGuard('ADMIN');
        const result = TestBed.runInInjectionContext(() =>
            guard(dummyRoute, dummyState)
        );

        expect(result).toBeFalse();
        expect(router.navigate).toHaveBeenCalledWith(['/login']);
    });
});
