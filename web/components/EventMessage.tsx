import { useDataChannel } from '@lizunlong/livekit-components-react';
import { useCallback, useState } from 'react';
import { EventPacket, Packet, PacketType } from '../lib/packet';
import type { ReceivedDataMessage } from '@livekit/components-core';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from './ui/alert-dialog';

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

  return <CloseRoomAlertDialog open={isCloseRoomAlertVisible} />;
};

function CloseRoomAlertDialog({ open }: { open: boolean }) {
  function handleClick() {
    window.close();
  }

  return (
    <AlertDialog open={open}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Video Interaction Over</AlertDialogTitle>
          <AlertDialogDescription>
            Video interaction is over, click the button below to leave the room
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogAction onClick={handleClick}>Leave</AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
