export interface CreateResponse {
  url: string;
  short_url: string;
  expiry: number;
}

export interface LinksResponse {
  count: number;
  links: string[];
}