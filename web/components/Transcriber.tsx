import { Box, Text } from '@chakra-ui/react';
import { useDataChannel } from '@lizunlong/livekit-components-react';
import { useCallback, useEffect, useState } from 'react';
import { GPTState, Packet, PacketType, StatePacket, TranscriptPacket } from '../lib/packet';
import type { ReceivedDataMessage } from '@livekit/components-core';

export const Transcriber = () => {
  const [visible, setVisible] = useState<boolean>(false);
  const [activity, setActivity] = useState<number>(Date.now());
  const [state, setState] = useState<GPTState>(GPTState.Idle);
  const [transcripts, setTranscripts] = useState<Map<string, string>>(new Map()); // transcription of every participant

  const onData = useCallback(
    (message: ReceivedDataMessage) => {
      const decoder = new TextDecoder();
      const packet = JSON.parse(decoder.decode(message.payload)) as Packet;
      console.log('packet:', packet);
      if (packet.type == PacketType.Transcript) {
        const transcript = packet.data as TranscriptPacket;
        const sid = transcript.sid;
        const text = transcript.name + ': ' + transcript.text;
        setTranscripts(new Map(transcripts.set(sid, text)));
        setActivity(Date.now());

        setVisible(true);
        // if (state == GPTState.Active) {
        //   setVisible(true);
        // }
      } else if (packet.type == PacketType.State) {
        const statePacket = packet.data as StatePacket;
        setState(statePacket.state);
      } else if (packet.type === PacketType.EventStopRoom) {
        console.log('data:', packet.data);
      }
    },
    [state],
  );

  useDataChannel(undefined, onData);

  useEffect(() => {
    const currentActivity = activity;
    const timeout = setTimeout(() => {
      if (currentActivity == activity) {
        setVisible(false);
        setTranscripts(new Map());
      }
    }, 5000);

    return () => clearTimeout(timeout);
  }, [activity]);

  return visible ? (
    <Box
      position="fixed"
      left="50%"
      transform="translateX(-50%)"
      paddingX="4px"
      bottom="8rem"
      bgColor="rgba(0, 0, 0, 0.6)"
    >
      {Array.from(transcripts.entries()).map((entry) => {
        const [key, value] = entry;
        return (
          <Text key={key} margin={0} fontSize={18}>
            {value}
          </Text>
        );
      })}
    </Box>
  ) : (
    <> </>
  );
};
