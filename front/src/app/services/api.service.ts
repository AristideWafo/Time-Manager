import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({ providedIn: 'root' })
export class ApiService {
    private baseUrl = '/api';

    constructor(private http: HttpClient) { }

    // Authentification
    login(credentials: { Email: string; Password: string }) {
        return this.http.post<{ token: string }>(
            `${this.baseUrl}/authentification`,
            credentials
        );
    }

    getUser() {
        return this.http.get<any>(`${this.baseUrl}/user`);
    }

    getAllUsers() {
        return this.http.get<any[]>('/api/users');
    }

    updateUser(data: any) {
        return this.http.post(`${this.baseUrl}/user/update`, data);
    }

    deleteUser() {
        return this.http.delete(`${this.baseUrl}/user`);
    }

    createUser(data: any) {
        return this.http.post(`${this.baseUrl}/user/create`, data);
    }

    // Presence
    postPresence(data: { Type: string; Timestamp: string }) {
        return this.http.post(`${this.baseUrl}/presence/create`, data);
    }

    getPresence() {
        return this.http.get(`${this.baseUrl}/presence`);
    }

    // TEAMS  

    getTeamByName(name: string) {
        return this.http.get(`${this.baseUrl}/team/${name}`);
    }

    createTeam(data: any) {
        return this.http.post(`${this.baseUrl}/team/create`, data);
    }

    updateTeam(data: any) {
        return this.http.put(`${this.baseUrl}/team/update`, data);
    }

    getTeamUsers(name: string) {
        return this.http.get<any[]>(`${this.baseUrl}/team/users/${name}`);
    }

    getAllTeams() {
        return this.http.get<any[]>(`${this.baseUrl}/team`);
    }

}
