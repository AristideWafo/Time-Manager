import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({ providedIn: 'root' })
export class ApiService {
    private baseUrl = 'http://localhost:8080';

    constructor(private http: HttpClient) { }

    // Authentification
    login(credentials: { email: string; password: string }) {
        return this.http.post<{ token: string }>(`${this.baseUrl}/authentification`, credentials);
    }

    // Utilisateur
    getUser(id: string) {
        return this.http.get(`${this.baseUrl}/user/${id}`);
    }

    updateUser(id: string, data: any) {
        return this.http.put(`${this.baseUrl}/user/${id}`, data);
    }

    deleteUser(id: string) {
        return this.http.delete(`${this.baseUrl}/user/${id}`);
    }

    // Équipes
    getAllTeams() {
        return this.http.get<any[]>(`${this.baseUrl}/team`);
    }

    createTeam(data: any) {
        return this.http.post(`${this.baseUrl}/team`, data);
    }

    updateTeam(id: string, data: any) {
        return this.http.put(`${this.baseUrl}/team/${id}`, data);
    }

    deleteTeam(id: string) {
        return this.http.delete(`${this.baseUrl}/team/${id}`);
    }

    // Pointage
    postPointage(data: { userId: string; isEntry: boolean }) {
        return this.http.post(`${this.baseUrl}/pointage`, data);
    }
}
