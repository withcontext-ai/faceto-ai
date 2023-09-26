import { TrackReference, TrackReferenceOrPlaceholder } from '@livekit/components-core';
import React from 'react';
import { LayoutContextType, useEnsureLayoutContext } from '@lizunlong/livekit-components-react';

export function usePinnedTracks(layoutContext?: LayoutContextType): TrackReference[] {
  layoutContext = useEnsureLayoutContext(layoutContext);
  return React.useMemo(() => {
    if (layoutContext?.pin.state !== undefined && layoutContext.pin.state.length >= 1) {
      // 添加类型检查
      return layoutContext.pin.state.filter(
        (track: TrackReferenceOrPlaceholder): track is TrackReference =>
          track.publication !== undefined,
      );
    }
    return [];
  }, [layoutContext]);
}
