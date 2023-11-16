import { useDataChannel } from '@lizunlong/livekit-components-react';
import { useCallback, useState } from 'react';
import { EventPacket, Packet, PacketType } from '../lib/packet';
import type { ReceivedDataMessage } from '@livekit/components-core';

export const EventMessage = () => {
  const [isCloseRoomAlertVisible, setIsCloseRoomAlertVisible] = useState(false);

  const onData = useCallback((message: ReceivedDataMessage) => {
    const decoder = new TextDecoder();
    const packet = JSON.parse(decoder.decode(message.payload)) as Packet;
    if (packet.type === PacketType.Event) {
      const eventPacket = packet.data as EventPacket;
      if (eventPacket.event === 'CloseRoom') {
        setIsCloseRoomAlertVisible(true);
      }
    }
  }, []);

  useDataChannel(undefined, onData);

  return null
};
