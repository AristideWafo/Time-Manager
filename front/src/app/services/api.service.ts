import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({ providedIn: 'root' })
export class ApiService {
    private baseUrl = '/api';

    constructor(private http: HttpClient) { }

    // Authentification
    login(credentials: { Email: string; Password: string }) {
        return this.http.post<{ token: string }>(
            `${this.baseUrl}/authentification`, credentials
        );
    }

    // Utilisateur connecté via token
    getUser() {
        return this.http.get<any>(`${this.baseUrl}/user`);
    }

    getAllUsers() {
        return this.http.get<any[]>(`${this.baseUrl}/users`);
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

    // Équipes (TEAM)

    getTeamByName(name: string) {
        return this.http.get<any>(`${this.baseUrl}/getTeamByName?name=${name}`);
    }

    updateTeam(name: string, data: any) {
        return this.http.put(`${this.baseUrl}/updateTeam?name=${name}`, data);
    }

    getTeamUsers(name: string) {
        return this.http.get<any[]>(`${this.baseUrl}/getTeamUsers?name=${name}`);
    }


    postPresence(data: { Type: string; Timestamp: string }) {
        return this.http.post(`${this.baseUrl}/presence/create`, data);
    }

    getPresence() {
        return this.http.get(`${this.baseUrl}/presence`);
    }
}
