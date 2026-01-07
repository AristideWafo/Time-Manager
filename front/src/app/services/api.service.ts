import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class ApiService {

    private baseUrl = '/api';

    constructor(private http: HttpClient) { }

    /* ===================== AUTH ===================== */

    login(credentials: { Email: string; Password: string }): Observable<any> {
        return this.http.post(
            `${this.baseUrl}/authentification`,
            credentials
        );
    }

    /* ===================== USER ===================== */

    getUser(): Observable<any> {
        return this.http.get(`${this.baseUrl}/user`);
    }

    updateUser(data: any): Observable<any> {
        return this.http.post(`${this.baseUrl}/user/update`, data);
    }

    /* ===================== ADMIN - USERS ===================== */

    getAllUsers(): Observable<any[]> {
        return this.http.get<any[]>(`${this.baseUrl}/admin/all/user`);
    }

    createUser(data: any): Observable<any> {
        return this.http.post(`${this.baseUrl}/admin/user/create`, data);
    }

    getUserById(id: string): Observable<any> {
        return this.http.get(`${this.baseUrl}/admin/${id}`);
    }

    updateUserTeam(id: string, data: any): Observable<any> {
        return this.http.put(`${this.baseUrl}/admin/update/team/user/${id}`, data);
    }

    /* ===================== ADMIN - TEAMS ===================== */

    getAllTeams(): Observable<any[]> {
        return this.http.get<any[]>(`${this.baseUrl}/admin/all/team`);
    }

    createTeam(data: any): Observable<any> {
        return this.http.post(`${this.baseUrl}/admin/team/create`, data);
    }

    updateTeam(data: any): Observable<any> {
        return this.http.put(`${this.baseUrl}/admin/team/update`, data);
    }

    getTeamUsers(name: string): Observable<any[]> {
        return this.http.get<any[]>(`${this.baseUrl}/admin/team/users/${name}`);
    }

    /* ===================== MANAGER ===================== */

    getManagerTeamUsers(): Observable<any[]> {
        return this.http.get<any[]>(
            `${this.baseUrl}/manager/team/users`
        );
    }

    getUserPresencesForManager(id: string): Observable<any[]> {
        return this.http.get<any[]>(
            `${this.baseUrl}/manager/team/user/presence/${id}`
        );
    }


    /* ===================== PRESENCE ===================== */

    getPresence(): Observable<any[]> {
        return this.http.get<any[]>(`${this.baseUrl}/presence`);
    }

    postPresence(data: { Type: string; Timestamp: string }): Observable<any> {
        return this.http.post(`${this.baseUrl}/presence/create`, data);
    }

    /* ===================== TEAM ===================== */

    getTeam(): Observable<any> {
        return this.http.get(`${this.baseUrl}/team/name`);
    }
}
