export enum PacketType {
  Transcript = 0,
  State,
  Error,
  Event,
}

export enum GPTState {
  Idle = 0,
  Loading,
  Speaking,
  Active,
}

export interface Packet {
  type: PacketType;
  data: TranscriptPacket | StatePacket | ErrorPacket | EventPacket;
}

export interface TranscriptPacket {
  sid: string;
  name: string;
  text: string;
  isFinal: boolean;
}

export interface StatePacket {
  state: GPTState;
}

export interface ErrorPacket {
  message: string;
}

export interface EventPacket {
  event: 'CloseRoom';
  name: string;
  sid: string;
  text: string;
}
