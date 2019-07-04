import PushStream from 'external/PushStream'

const newPushStreamInstance = (settings) => new PushStream(settings)

export default {
  newPushStreamInstance,
}
