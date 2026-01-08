import { ComponentFixture, TestBed } from '@angular/core/testing';
import { LoginComponent } from '../../view/pages/login/login.component';
import { ApiService } from '../../services/api.service';
import { Router } from '@angular/router';
import { of, throwError } from 'rxjs';
import { FormsModule } from '@angular/forms';
import { RouterTestingModule } from '@angular/router/testing';

describe('LoginComponent', () => {
    let component: LoginComponent;
    let fixture: ComponentFixture<LoginComponent>;
    let apiServiceSpy: jasmine.SpyObj<ApiService>;
    let router: Router;

    beforeEach(async () => {
        apiServiceSpy = jasmine.createSpyObj('ApiService', ['login']);

        await TestBed.configureTestingModule({
            imports: [
                FormsModule,
                RouterTestingModule
            ],
            providers: [
                { provide: ApiService, useValue: apiServiceSpy }
            ]
        }).compileComponents();

        fixture = TestBed.createComponent(LoginComponent);
        component = fixture.componentInstance;
        router = TestBed.inject(Router);

        spyOn(router, 'navigate');
        spyOn(localStorage, 'setItem');
    });

    it('should create the component', () => {
        expect(component).toBeTruthy();
    });

    it('should call api.login with email and password', () => {
        apiServiceSpy.login.and.returnValue(of({
            Token: 'token',
            _id: '123',
            Role: 'EMPLOYEE'
        }));

        component.email = 'test@mail.com';
        component.password = 'password';
        component.onLogin();

        expect(apiServiceSpy.login).toHaveBeenCalledWith({
            Email: 'test@mail.com',
            Password: 'password'
        });
    });

    it('should store data and redirect to admin home', () => {
        apiServiceSpy.login.and.returnValue(of({
            Token: 'token',
            _id: '123',
            Role: 'ADMIN',
            FirstName: 'John',
            LastName: 'Doe',
            Team: 'DEV'
        }));

        component.onLogin();

        expect(localStorage.setItem).toHaveBeenCalledWith('access_token', 'token');
        expect(router.navigate).toHaveBeenCalledWith(['/admin-home']);
    });

    it('should display error message on login failure', () => {
        apiServiceSpy.login.and.returnValue(
            throwError(() => ({
                error: { message: 'Invalid credentials' }
            }))
        );

        component.onLogin();

        expect(component.error).toBe('Invalid credentials');
        expect(component.loading).toBeFalse();
    });

    it('should stop if already loading', () => {
        component.loading = true;
        component.onLogin();
        expect(apiServiceSpy.login).not.toHaveBeenCalled();
    });
});
