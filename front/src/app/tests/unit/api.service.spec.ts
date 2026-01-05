import { TestBed } from '@angular/core/testing';
import { ApiService } from '../../services/api.service';
import {
    HttpClientTestingModule,
    HttpTestingController
} from '@angular/common/http/testing';

describe('ApiService', () => {
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

    /* ===================== AUTH ===================== */

    it('should call POST /api/authentification on login', () => {
        const credentials = {
            Email: 'test@mail.com',
            Password: 'password'
        };

        service.login(credentials).subscribe();

        const req = httpMock.expectOne('/api/authentification');
        expect(req.request.method).toBe('POST');
        expect(req.request.body).toEqual(credentials);

        req.flush({ Token: 'token', _id: '123' });
    });

    /* ===================== USER ===================== */

    it('should call GET /api/user', () => {
        service.getUser().subscribe();

        const req = httpMock.expectOne('/api/user');
        expect(req.request.method).toBe('GET');

        req.flush({});
    });

    /* ===================== ADMIN USERS ===================== */

    it('should call GET /api/admin/all/user', () => {
        service.getAllUsers().subscribe();

        const req = httpMock.expectOne('/api/admin/all/user');
        expect(req.request.method).toBe('GET');

        req.flush([]);
    });

    it('should call GET /api/admin/{id}', () => {
        service.getUserById('42').subscribe();

        const req = httpMock.expectOne('/api/admin/42');
        expect(req.request.method).toBe('GET');

        req.flush({});
    });

    /* ===================== ADMIN TEAMS ===================== */

    it('should call GET /api/admin/all/team', () => {
        service.getAllTeams().subscribe();

        const req = httpMock.expectOne('/api/admin/all/team');
        expect(req.request.method).toBe('GET');

        req.flush([]);
    });

    /* ===================== MANAGER ===================== */

    it('should call GET /api/manager/team/users/{name} with Authorization header', () => {
        const teamName = 'DEV';
        localStorage.setItem('token', 'fake-token');

        service.getManagerTeamUsersByName(teamName).subscribe();

        const req = httpMock.expectOne(`/api/manager/team/users/${teamName}`);
        expect(req.request.method).toBe('GET');
        expect(req.request.headers.get('Authorization')).toBe('Bearer fake-token');

        req.flush([]);
    });

    /* ===================== PRESENCE ===================== */

    it('should call POST /api/presence/create', () => {
        const data = { Type: 'IN', Timestamp: '2026-01-01T08:00:00' };

        service.postPresence(data).subscribe();

        const req = httpMock.expectOne('/api/presence/create');
        expect(req.request.method).toBe('POST');
        expect(req.request.body).toEqual(data);

        req.flush({});
    });
});
