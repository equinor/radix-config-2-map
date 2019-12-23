# radix-config-2-map

The radix-config-2-map gives radix-pipeline a way to obtain cloned radix config from a config branch and persist into a config map of the app namespace, in order for it to be validated and applied by the pipeline.

Build is done using Github actions. There are secrets defined for the actions to be able to push to radixdev, radixprod and radixus. These are the corresponding credentials for radix-cr-cicd-dev and radix-cr-cicd-prod service accounts

[![Build Status](https://github.com/equinor/radix-config-2-map/workflows/radix-config-2-map-build/badge.svg)](https://github.com/equinor/radix-config-2-map/actions?query=workflow%3Aradix-config-2-map-build)

.
