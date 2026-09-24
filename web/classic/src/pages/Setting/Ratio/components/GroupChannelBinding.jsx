/*
Copyright (C) 2025 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/
import React, { useCallback, useEffect, useMemo, useState } from 'react';
import {
  Button,
  Select,
  Space,
  Spin,
  Table,
  Tag,
  Typography,
} from '@douyinfe/semi-ui';
import { useTranslation } from 'react-i18next';
import { API, showError, showSuccess } from '../../../helpers';

const { Text } = Typography;

const CHANNEL_STATUS_ENABLED = 1;
const WILDCARD_MODEL = '*';

function pinsForGroup(groupChannels, group) {
  return groupChannels?.[group] || {};
}

function preferredModel(models, pins) {
  const pinned = Object.keys(pins || {}).sort((left, right) => {
    if (left === WILDCARD_MODEL) return 1;
    if (right === WILDCARD_MODEL) return -1;
    return left.localeCompare(right);
  });
  const fromPins = pinned.find((name) => models.includes(name));
  if (fromPins) return fromPins;
  return models.find((name) => name !== WILDCARD_MODEL) || models[0] || '';
}

function channelSupportsModel(channel, model) {
  if (model === WILDCARD_MODEL) return true;
  return (channel.models || []).includes(model);
}

export default function GroupChannelBinding() {
  const { t } = useTranslation();
  const [loading, setLoading] = useState(false);
  const [saving, setSaving] = useState(false);
  const [groups, setGroups] = useState([]);
  const [channels, setChannels] = useState([]);
  const [groupChannels, setGroupChannels] = useState({});
  const [groupModels, setGroupModels] = useState({});
  const [group, setGroup] = useState('');
  const [model, setModel] = useState('');
  const [selected, setSelected] = useState([]);

  const applyPayload = useCallback((data) => {
    if (!data) return null;
    setGroups(data.groups || []);
    setChannels(data.channels || []);
    setGroupChannels(data.group_channels || {});
    setGroupModels(data.group_models || {});
    return data;
  }, []);

  const modelNamesFor = useCallback((data, currentGroup) => {
    const names = new Set([WILDCARD_MODEL]);
    (data?.group_models?.[currentGroup] || []).forEach((name) => names.add(name));
    Object.keys(data?.group_channels?.[currentGroup] || {}).forEach((name) =>
      names.add(name),
    );
    (data?.channels || []).forEach((channel) => {
      (channel.models || []).forEach((name) => names.add(name));
    });
    return Array.from(names).sort((left, right) => {
      if (left === WILDCARD_MODEL) return -1;
      if (right === WILDCARD_MODEL) return 1;
      return left.localeCompare(right);
    });
  }, []);

  const load = useCallback(
    async (keepGroup, keepModel) => {
      setLoading(true);
      try {
        const res = await API.get('/api/group/channels');
        const data = applyPayload(res.data?.data);
        const currentGroup =
          keepGroup && data?.groups?.includes(group)
            ? group
            : data?.groups?.[0] || '';
        const models = modelNamesFor(data, currentGroup);
        const currentModel =
          keepModel && models.includes(model)
            ? model
            : preferredModel(
                models,
                pinsForGroup(data?.group_channels, currentGroup),
              );
        setGroup(currentGroup);
        setModel(currentModel);
        setSelected(
          pinsForGroup(data?.group_channels, currentGroup)[currentModel] || [],
        );
      } catch (error) {
        showError(t('加载分组渠道绑定失败'));
      } finally {
        setLoading(false);
      }
    },
    [applyPayload, group, model, modelNamesFor, t],
  );

  useEffect(() => {
    load(false, false);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleGroupChange = (value) => {
    const models = modelNamesFor(
      {
        group_models: groupModels,
        group_channels: groupChannels,
        channels,
      },
      value,
    );
    const nextModel = preferredModel(
      models,
      pinsForGroup(groupChannels, value),
    );
    setGroup(value);
    setModel(nextModel);
    setSelected(pinsForGroup(groupChannels, value)[nextModel] || []);
  };

  const handleModelChange = (value) => {
    setModel(value);
    setSelected(pinsForGroup(groupChannels, group)[value] || []);
  };

  const channelById = useMemo(
    () => new Map(channels.map((item) => [item.id, item])),
    [channels],
  );

  const matchingChannels = useMemo(() => {
    if (!model) return [];
    return channels.filter((channel) => channelSupportsModel(channel, model));
  }, [channels, model]);

  const currentPins = useMemo(() => {
    const pins = pinsForGroup(groupChannels, group);
    return Object.keys(pins)
      .sort((left, right) => {
        if (left === WILDCARD_MODEL) return -1;
        if (right === WILDCARD_MODEL) return 1;
        return left.localeCompare(right);
      })
      .map((name) => ({
        model: name,
        channelIds: pins[name] || [],
      }));
  }, [groupChannels, group]);

  const modelOptions = useMemo(
    () =>
      modelNamesFor(
        {
          group_models: groupModels,
          group_channels: groupChannels,
          channels,
        },
        group,
      ),
    [group, groupChannels, groupModels, modelNamesFor],
  );

  const persist = async (ids) => {
    if (!group || !model) return;
    setSaving(true);
    try {
      const res = await API.put('/api/group/channels', {
        group,
        model,
        channel_ids: ids,
      });
      if (!res.data?.success) {
        showError(res.data?.message || t('保存失败'));
        return;
      }
      showSuccess(
        ids.length === 0
          ? t('已取消该模型的渠道绑定')
          : t('分组渠道绑定已保存'),
      );
      await load(true, true);
    } catch (error) {
      showError(t('保存失败'));
    } finally {
      setSaving(false);
    }
  };

  const columns = [
    {
      title: t('渠道'),
      dataIndex: 'name',
      render: (text, record) => (
        <Text>
          {text} #{record.id}
        </Text>
      ),
    },
    {
      title: t('状态'),
      dataIndex: 'status',
      width: 110,
      render: (status) =>
        status === CHANNEL_STATUS_ENABLED ? (
          <Tag color='green'>{t('启用')}</Tag>
        ) : (
          <Tag color='grey'>{t('已禁用')}</Tag>
        ),
    },
  ];

  return (
    <Spin spinning={loading}>
      <Space vertical align='start' spacing={12} style={{ width: '100%' }}>
        <Text type='tertiary'>
          {t(
            '为用户分组中的某个模型指定渠道。该分组调用此模型时只会走绑定的渠道；未绑定的模型仍按渠道自身的分组设置选择。',
          )}
        </Text>
        <Space wrap>
          <Text strong>{t('用户分组')}</Text>
          <Select
            value={group}
            onChange={handleGroupChange}
            style={{ width: 200 }}
            placeholder={t('选择分组')}
          >
            {groups.map((name) => (
              <Select.Option key={name} value={name}>
                {name}
              </Select.Option>
            ))}
          </Select>
          <Text strong>{t('模型')}</Text>
          <Select
            value={model}
            onChange={handleModelChange}
            style={{ width: 240 }}
            placeholder={t('选择模型')}
            filter
          >
            {modelOptions.map((name) => (
              <Select.Option key={name} value={name}>
                {name === WILDCARD_MODEL ? t('全部模型') : name}
              </Select.Option>
            ))}
          </Select>
        </Space>
        {currentPins.length > 0 && (
          <div>
            <Text strong>{t('当前绑定')}</Text>
            <div style={{ marginTop: 6 }}>
              {currentPins.map((pin) => (
                <Tag
                  key={pin.model}
                  color={pin.model === model ? 'blue' : 'white'}
                  style={{ marginRight: 6, marginBottom: 6 }}
                >
                  {pin.model === WILDCARD_MODEL ? t('全部模型') : pin.model}
                  {' → '}
                  {pin.channelIds
                    .map((id) => {
                      const channel = channelById.get(id);
                      return channel ? `${channel.name} #${id}` : `#${id}`;
                    })
                    .join(', ')}
                </Tag>
              ))}
            </div>
          </div>
        )}
        <Table
          rowKey='id'
          columns={columns}
          dataSource={matchingChannels}
          pagination={{ pageSize: 8 }}
          rowSelection={{
            selectedRowKeys: selected,
            onChange: (keys) => setSelected(keys),
            getCheckboxProps: (record) => ({
              disabled: record.status !== CHANNEL_STATUS_ENABLED,
            }),
          }}
        />
        <Space>
          <Button
            theme='solid'
            loading={saving}
            disabled={!group || !model}
            onClick={() => persist(selected)}
          >
            {t('保存绑定')}
          </Button>
          <Button
            loading={saving}
            disabled={!group || !model}
            onClick={() => persist([])}
          >
            {t('取消绑定')}
          </Button>
          {selected.length === 0 && (
            <Text type='tertiary' size='small'>
              {t('未选择渠道，保存后将取消该模型的绑定')}
            </Text>
          )}
        </Space>
      </Space>
    </Spin>
  );
}
