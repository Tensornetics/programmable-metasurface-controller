#include <iostream>
#include <unordered_map>

#include "brainwave_signal.h"
#include "metasurface.h"

// Maps brainwave attention intensity values to scattering patterns
const std::unordered_map<int, ScatteringPattern> attentionIntensityMap = {
    {1, ScatteringPattern::SPATIAL_HOLOGRAM},
    {2, ScatteringPattern::BEAM_STEERING},
    {3, ScatteringPattern::FOCUSED_BEAM},
    {4, ScatteringPattern::WIDE_BEAM},
};

void controlMetasurface(const BrainwaveSignal& signal, Metasurface& surface) {
    int attentionIntensity = getAttentionIntensity(signal);
    ScatteringPattern pattern = attentionIntensityMap.at(attentionIntensity);
    surface.setScatteringPattern(pattern);
}
