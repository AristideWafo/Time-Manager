import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { CreateUserComponent } from '../../view/pages/create-user/create_user.component';
import { Router } from '@angular/router';
import { ApiService } from '../../services/api.service';
import { of, throwError } from 'rxjs';

describe('CreateUserComponent', () => {
    let component: CreateUserComponent;
    let fixture: ComponentFixture<CreateUserComponent>;
    let apiSpy: jasmine.SpyObj<ApiService>;
    let routerSpy: jasmine.SpyObj<Router>;

    beforeEach(() => {
        apiSpy = jasmine.createSpyObj('ApiService', ['getAllTeams', 'createUser', 'updateUserTeam']);
        routerSpy = jasmine.createSpyObj('Router', ['navigate']);

        TestBed.configureTestingModule({
            imports: [
                FormsModule,
                CreateUserComponent // ✅ standalone component
            ],
            providers: [
                { provide: ApiService, useValue: apiSpy },
                { provide: Router, useValue: routerSpy }
            ]
        }).compileComponents();

        fixture = TestBed.createComponent(CreateUserComponent);
        component = fixture.componentInstance;
    });

    it('should load teams on init', () => {
        const fakeTeams = [{ _id: '1', Name: 'DEV' }];

        apiSpy.getAllTeams.and.returnValue(
            of({ teams: fakeTeams } as any)
        );

        component.ngOnInit();

        expect(apiSpy.getAllTeams).toHaveBeenCalled();
        expect(component.teams).toEqual(fakeTeams);
    });


    it('should create user and assign team successfully', () => {
        const createdUser = { user: { _id: '123' } };
        apiSpy.createUser.and.returnValue(of(createdUser));
        apiSpy.updateUserTeam.and.returnValue(of({}));

        component.firstName = 'John';
        component.lastName = 'Doe';
        component.email = 'john@mail.com';
        component.password = '123456';
        component.role = 'EMPLOYEE';
        component.teamId = '1';

        component.createUser();

        expect(apiSpy.createUser).toHaveBeenCalledWith({
            FirstName: 'John',
            LastName: 'Doe',
            Email: 'john@mail.com',
            Password: '123456',
            Role: 'EMPLOYEE'
        });

        expect(apiSpy.updateUserTeam).toHaveBeenCalledWith('123', { teamId: '1' });
        expect(component.message).toBe('Utilisateur créé et assigné à une équipe ✅');
    });

    it('should handle error on createUser', () => {
        apiSpy.createUser.and.returnValue(throwError(() => ({ error: 'fail' })));

        component.createUser();

        expect(component.message).toBe('Erreur lors de la création de l’utilisateur ❌');
    });

    it('should handle error on updateUserTeam', () => {
        const createdUser = { user: { _id: '123' } };
        apiSpy.createUser.and.returnValue(of(createdUser));
        apiSpy.updateUserTeam.and.returnValue(throwError(() => ({ error: 'fail' })));
        component.teamId = '1';

        component.createUser();

        expect(component.message).toBe('Utilisateur créé mais erreur lors de l’assignation de la team ❌');
    });
});
