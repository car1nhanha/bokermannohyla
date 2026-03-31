export interface ResponseCreateLink {
  Id: string;
  Ip: string;
  CreatedAt: number;
  Accesses: number;
  Original: string;
}

export interface ResponseGetLink {
  Id: string;
  Ip: string;
  CreatedAt: number;
  Accesses: number;
  Original: string;
}

export class ApiService {
  public static async getLinks(): Promise<ResponseGetLink[]> {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/links`);
    if (!response.ok) {
      throw new Error("Failed to fetch links");
    }

    return response.json();
  }

  public static async createLink(link: string): Promise<ResponseCreateLink> {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/link`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        original: link,
      }),
    });
    if (!response.ok) {
      throw new Error("Failed to create link");
    }
    return response.json();
  }

  public static async deleteLink(id: string): Promise<void> {
    const response = await fetch(`${import.meta.env.VITE_API_URL}/link/${id}`, {
      method: "DELETE",
    });
    if (!response.ok) {
      throw new Error("Failed to delete link");
    }
  }
}
