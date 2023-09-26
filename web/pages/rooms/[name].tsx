import { useMemo, useRef, useState } from 'react';
import Head from 'next/head';
import type { NextPage } from 'next';
import { useRouter } from 'next/router';
import { LiveKitRoom, useToken, formatChatMessageLinks } from '@lizunlong/livekit-components-react';
import { RoomOptions, VideoPresets } from 'livekit-client';
import { toast } from 'react-hot-toast';
import { decode } from 'js-base64';
import { useServerUrl } from '../../lib/client-utils';
import { VideoConference } from '../../components/VideoConference';
import { LocalUserChoices, PreJoin } from '../../components/PreJoin';
import useCheckRoom from '../../hooks/useCheckRoom';

const nameReg = new RegExp('^[a-zA-Z0-9_-]{1,64}$');

function getValueOfConfig(config?: string, key?: string) {
  if (config && typeof config === 'string') {
    try {
      const str = decode(config)
      const obj = JSON.parse(str);
      if (key && obj[key]) {
        return obj[key];
      }
    } catch (error) {
      return undefined
    }
  }
  return undefined;
}

const Home: NextPage = () => {
  const router = useRouter();
  const roomName =  router.query.name as string;
  const config = router.query.c as string | undefined;
  const { checkRoom } = useCheckRoom();

  const defaultUsername = useMemo(() => getValueOfConfig(config, 'username'), [config])
  const botName = useMemo(() => getValueOfConfig(config, 'botname'), [config])

  const [preJoinChoices, setPreJoinChoices] = useState<LocalUserChoices | undefined>(undefined);
  return (
    <>
      <Head>
        <title>LiveKit Meet</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main data-lk-theme="default">
        {roomName && !Array.isArray(roomName) && preJoinChoices ? (
          <ActiveRoom
            roomName={roomName}
            userChoices={preJoinChoices}
            onLeave={() => {
              setPreJoinChoices(undefined);
            }}
            botName={botName}
          ></ActiveRoom>
        ) : (
          <div style={{ display: 'grid', placeItems: 'center', height: '100%' }}>
            <PreJoin
              onError={(err) => console.log('error while setting up prejoin', err)}
              // onValidate={(values) => {
              //   return nameReg.test(values.username);
              // }}
              defaults={{
                username: defaultUsername,
                videoEnabled: true,
                audioEnabled: true,
              }}
              onSubmit={async (values) => {
                console.log('Joining with: ', values);
                const isValid = await checkRoom(roomName, config);
                if (isValid) {
                  setPreJoinChoices(values);
                } else {
                  toast.error('The room is no longer available as the interview has ended.');
                }
              }}
            ></PreJoin>
          </div>
        )}
      </main>
    </>
  );
};

export default Home;

type ActiveRoomProps = {
  userChoices: LocalUserChoices;
  roomName: string;
  region?: string;
  onLeave?: () => void;
  botName?: string;
};

const ActiveRoom = ({ roomName, userChoices, onLeave, botName }: ActiveRoomProps) => {
  const identity = useRef(`${Date.now()}`)
  const token = useToken("/api/token", roomName, {
  // const token = useToken(process.env.NEXT_PUBLIC_LK_TOKEN_ENDPOINT, roomName, {
    userInfo: {
      identity: identity.current,
      name: userChoices.username,
      metadata: JSON.stringify({ languageCode: userChoices.language }),
    },
  });

  const router = useRouter();
  const { region, hq } = router.query;

  const liveKitUrl = useServerUrl(region as string | undefined);

  const roomOptions = useMemo((): RoomOptions => {
    return {
      videoCaptureDefaults: {
        deviceId: userChoices.videoDeviceId ?? undefined,
        resolution: hq === 'true' ? VideoPresets.h2160 : VideoPresets.h720,
      },
      publishDefaults: {
        red: false,
        dtx: false,
        videoSimulcastLayers:
          hq === 'true'
            ? [VideoPresets.h1080, VideoPresets.h720]
            : [VideoPresets.h540, VideoPresets.h216],
      },
      audioCaptureDefaults: {
        deviceId: userChoices.audioDeviceId ?? undefined,
      },
      adaptiveStream: { pixelDensity: 'screen' },
      dynacast: true,
    };
  }, [userChoices, hq]);

  return (
    <>
      {liveKitUrl && (
        <LiveKitRoom
          token={token}
          serverUrl={liveKitUrl}
          options={roomOptions}
          video={userChoices.videoEnabled}
          audio={userChoices.audioEnabled}
          onDisconnected={onLeave}
        >
          <VideoConference chatMessageFormatter={formatChatMessageLinks} botName={botName} />
        </LiveKitRoom>
      )}
    </>
  );
};
