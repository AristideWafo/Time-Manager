import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({ providedIn: 'root' })
export class ApiService {

    private baseUrl = 'http://localhost:8080';

    constructor(private http: HttpClient) { }

    login(credentials: { email: string; password: string }) {
        return this.http.post<{ token: string }>(
            `${this.baseUrl}/authentification`,
            credentials
        );
    }

    getUser(id: number) {
        return this.http.get(`${this.baseUrl}/user/${id}`);
    }

    postUser(data: any) {
        return this.http.post(`${this.baseUrl}/user/`, data);
    }

    updateUser(id: number, data: any) {
        return this.http.put(`${this.baseUrl}/user/${id}`, data);
    }

    deleteUser(id: number) {
        return this.http.delete(`${this.baseUrl}/user/${id}`);
    }

    getAllTeams() {
        return this.http.get<any[]>(`/team`);
    }

    createTeam(data: any) {
        return this.http.post(`/team`, data);
    }

    updateTeam(id: string, data: any) {
        return this.http.put(`/team/${id}`, data);
    }

    deleteTeam(id: string) {
        return this.http.delete(`/team/${id}`);
    }


    updatePointage(id: number, data: boolean) {
        return this.http.put(`${this.baseUrl}/pointage/${id}`, data);
    }
}
