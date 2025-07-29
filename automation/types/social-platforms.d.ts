// Shared TypeScript types/interfaces for social platform automation

export interface SocialAccount {
  platform: string;
  username: string;
  password: string;
  cookiesPath?: string;
}

export interface AutomationTask {
  platform: string;
  action: string;
  content?: string;
  target?: string;
  schedule?: string;
}
