import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({ providedIn: 'root' })
export class ApiService {

    private baseUrl = '/api';

    constructor(private http: HttpClient) { }

    /* ===================== AUTH ===================== */

    login(credentials: { Email: string; Password: string }) {
        return this.http.post(
            `${this.baseUrl}/authentification`,
            credentials
        );
    }

    /* ===================== USER ===================== */

    // GET /api/user
    getUser() {
        return this.http.get(`${this.baseUrl}/user`);
    }

    // POST /api/user/update
    updateUser(data: any) {
        return this.http.post(`${this.baseUrl}/user/update`, data);
    }

    /* ===================== ADMIN - USERS ===================== */

    // GET /api/admin/all/user
    getAllUsers() {
        return this.http.get<any[]>(
            `${this.baseUrl}/admin/all/user`
        );
    }

    // POST /api/admin/user/create
    createUser(data: any) {
        return this.http.post(
            `${this.baseUrl}/admin/user/create`,
            data
        );
    }

    // GET /api/admin/{id}
    getUserById(id: string) {
        return this.http.get(
            `${this.baseUrl}/admin/${id}`
        );
    }

    // PUT /api/admin/update/team/user/{id}
    updateUserTeam(id: string, data: any) {
        return this.http.put(
            `${this.baseUrl}/admin/update/team/user/${id}`,
            data
        );
    }

    /* ===================== ADMIN - TEAMS ===================== */

    // GET /api/admin/all/team
    getAllTeams() {
        return this.http.get<any[]>(
            `${this.baseUrl}/admin/all/team`
        );
    }

    // POST /api/admin/team/create
    createTeam(data: any) {
        return this.http.post(
            `${this.baseUrl}/admin/team/create`,
            data
        );
    }

    // PUT /api/admin/team/update
    updateTeam(data: any) {
        return this.http.put(
            `${this.baseUrl}/admin/team/update`,
            data
        );
    }

    // GET /api/admin/team/users/{name}
    getTeamUsers(name: string) {
        return this.http.get<any[]>(
            `${this.baseUrl}/admin/team/users/${name}`
        );
    }

    /* ===================== MANAGER ===================== */

    // GET /api/manager/team/users/{name}
    getManagerTeamUsers() {
        return this.http.get<any[]>(
            `${this.baseUrl}/manager/team/users/{name}`
        );
    }

    // GET /api/manager/team/user/presence/{id}
    getUserPresencesForManager(id: string) {
        return this.http.get<any[]>(
            `${this.baseUrl}/manager/team/user/presence/${id}`
        );
    }

    /* ===================== PRESENCE ===================== */

    // GET /api/presence
    getPresence() {
        return this.http.get<any[]>(
            `${this.baseUrl}/presence`
        );
    }

    // POST /api/presence/create
    postPresence(data: { Type: string; Timestamp: string }) {
        return this.http.post(
            `${this.baseUrl}/presence/create`,
            data
        );
    }

    /* ===================== TEAM ===================== */

    // GET /api/team/name
    getTeam() {
        return this.http.get(
            `${this.baseUrl}/team/name`
        );
    }
}
