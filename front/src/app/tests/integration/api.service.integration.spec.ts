// api.service.integration.spec.ts
import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { ApiService } from '../../services/api.service';

describe('ApiService Integration', () => {
    let service: ApiService;
    let httpMock: HttpTestingController;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
            providers: [ApiService]
        });

        service = TestBed.inject(ApiService);
        httpMock = TestBed.inject(HttpTestingController);
    });

    afterEach(() => {
        httpMock.verify();
    });

    it('should fetch all teams', () => {
        const mockResponse = { teams: [{ _id: '1', Name: 'DEV' }] };

        service.getAllTeams().subscribe((res: any) => {
            expect(res.teams.length).toBe(1);
            expect(res.teams[0]._id).toBe('1');
            expect(res.teams[0].Name).toBe('DEV');
        });

        const req = httpMock.expectOne('/api/admin/all/team');
        expect(req.request.method).toBe('GET');
        req.flush(mockResponse);
    });

    it('should create a user', () => {
        const userPayload = {
            FirstName: 'John',
            LastName: 'Doe',
            Email: 'john@mail.com',
            Password: '123456',
            Role: 'EMPLOYEE'
        };

        const mockResponse = { user: { _id: '123' } };

        service.createUser(userPayload).subscribe((res: any) => {
            expect(res.user._id).toBe('123');
        });

        const req = httpMock.expectOne('/api/admin/user/create');
        expect(req.request.method).toBe('POST');
        expect(req.request.body).toEqual(userPayload);
        req.flush(mockResponse);
    });

    it('should assign a team to a user', () => {
        const mockResponse = { success: true };

        service.updateUserTeam('123', { teamId: '1' }).subscribe((res: any) => {
            expect(res.success).toBeTrue();
        });

        const req = httpMock.expectOne('/api/admin/update/team/user/123');
        expect(req.request.method).toBe('PUT');
        expect(req.request.body).toEqual({ teamId: '1' });
        req.flush(mockResponse);
    });

    it('should handle API error gracefully', () => {
        service.createUser({} as any).subscribe({
            next: () => fail('Should have failed'),
            error: (error) => {
                expect(error.status).toBe(400);
            }
        });

        const req = httpMock.expectOne('/api/admin/user/create'); // URL corrigée
        req.flush({ message: 'Invalid data' }, { status: 400, statusText: 'Bad Request' });
    });
});
