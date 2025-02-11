import { required, minLength, maxLength, minValue, maxValue, helpers } from '@vuelidate/validators';

const workload_version_id = {
  required: helpers.withMessage('必須', required),
}

const workload_name = {
  required: helpers.withMessage('必須', required),
  minLength: helpers.withMessage('3文字以上', minLength(3)),
  maxLength: helpers.withMessage('20文字以下', maxLength(20)),
};

const workload_priority = {
  required: helpers.withMessage('必須', required),
  minValue: helpers.withMessage('0以上', minValue(0)),
  maxValue: helpers.withMessage('100,000以下', maxValue(100000)),
}

export const validationRules = {
  workload_name: workload_name,
  workload_priority: workload_priority,
  workload_version_id: workload_version_id,
};
