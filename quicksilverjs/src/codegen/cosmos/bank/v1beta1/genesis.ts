import { Params, ParamsAmino, ParamsSDKType, Metadata, MetadataAmino, MetadataSDKType } from "./bank";
import { Coin, CoinAmino, CoinSDKType } from "../../base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";
import { isSet, DeepPartial } from "../../../helpers";
export const protobufPackage = "cosmos.bank.v1beta1";
/** GenesisState defines the bank module's genesis state. */
export interface GenesisState {
  /** params defines all the paramaters of the module. */
  params: Params;
  /** balances is an array containing the balances of all the accounts. */
  balances: Balance[];
  /**
   * supply represents the total supply. If it is left empty, then supply will be calculated based on the provided
   * balances. Otherwise, it will be used to validate that the sum of the balances equals this amount.
   */
  supply: Coin[];
  /** denom_metadata defines the metadata of the differents coins. */
  denomMetadata: Metadata[];
}
export interface GenesisStateProtoMsg {
  typeUrl: "/cosmos.bank.v1beta1.GenesisState";
  value: Uint8Array;
}
/** GenesisState defines the bank module's genesis state. */
export interface GenesisStateAmino {
  /** params defines all the paramaters of the module. */
  params?: ParamsAmino;
  /** balances is an array containing the balances of all the accounts. */
  balances: BalanceAmino[];
  /**
   * supply represents the total supply. If it is left empty, then supply will be calculated based on the provided
   * balances. Otherwise, it will be used to validate that the sum of the balances equals this amount.
   */
  supply: CoinAmino[];
  /** denom_metadata defines the metadata of the differents coins. */
  denom_metadata: MetadataAmino[];
}
export interface GenesisStateAminoMsg {
  type: "cosmos-sdk/GenesisState";
  value: GenesisStateAmino;
}
/** GenesisState defines the bank module's genesis state. */
export interface GenesisStateSDKType {
  params: ParamsSDKType;
  balances: BalanceSDKType[];
  supply: CoinSDKType[];
  denom_metadata: MetadataSDKType[];
}
/**
 * Balance defines an account address and balance pair used in the bank module's
 * genesis state.
 */
export interface Balance {
  /** address is the address of the balance holder. */
  address: string;
  /** coins defines the different coins this balance holds. */
  coins: Coin[];
}
export interface BalanceProtoMsg {
  typeUrl: "/cosmos.bank.v1beta1.Balance";
  value: Uint8Array;
}
/**
 * Balance defines an account address and balance pair used in the bank module's
 * genesis state.
 */
export interface BalanceAmino {
  /** address is the address of the balance holder. */
  address: string;
  /** coins defines the different coins this balance holds. */
  coins: CoinAmino[];
}
export interface BalanceAminoMsg {
  type: "cosmos-sdk/Balance";
  value: BalanceAmino;
}
/**
 * Balance defines an account address and balance pair used in the bank module's
 * genesis state.
 */
export interface BalanceSDKType {
  address: string;
  coins: CoinSDKType[];
}
function createBaseGenesisState(): GenesisState {
  return {
    params: Params.fromPartial({}),
    balances: [],
    supply: [],
    denomMetadata: []
  };
}
export const GenesisState = {
  typeUrl: "/cosmos.bank.v1beta1.GenesisState",
  aminoType: "cosmos-sdk/GenesisState",
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.balances) {
      Balance.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.supply) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.denomMetadata) {
      Metadata.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.balances.push(Balance.decode(reader, reader.uint32()));
          break;
        case 3:
          message.supply.push(Coin.decode(reader, reader.uint32()));
          break;
        case 4:
          message.denomMetadata.push(Metadata.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): GenesisState {
    const obj = createBaseGenesisState();
    if (isSet(object.params)) obj.params = Params.fromJSON(object.params);
    if (Array.isArray(object?.balances)) obj.balances = object.balances.map((e: any) => Balance.fromJSON(e));
    if (Array.isArray(object?.supply)) obj.supply = object.supply.map((e: any) => Coin.fromJSON(e));
    if (Array.isArray(object?.denomMetadata)) obj.denomMetadata = object.denomMetadata.map((e: any) => Metadata.fromJSON(e));
    return obj;
  },
  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.balances) {
      obj.balances = message.balances.map(e => e ? Balance.toJSON(e) : undefined);
    } else {
      obj.balances = [];
    }
    if (message.supply) {
      obj.supply = message.supply.map(e => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.supply = [];
    }
    if (message.denomMetadata) {
      obj.denomMetadata = message.denomMetadata.map(e => e ? Metadata.toJSON(e) : undefined);
    } else {
      obj.denomMetadata = [];
    }
    return obj;
  },
  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = createBaseGenesisState();
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    }
    message.balances = object.balances?.map(e => Balance.fromPartial(e)) || [];
    message.supply = object.supply?.map(e => Coin.fromPartial(e)) || [];
    message.denomMetadata = object.denomMetadata?.map(e => Metadata.fromPartial(e)) || [];
    return message;
  },
  fromSDK(object: GenesisStateSDKType): GenesisState {
    return {
      params: object.params ? Params.fromSDK(object.params) : undefined,
      balances: Array.isArray(object?.balances) ? object.balances.map((e: any) => Balance.fromSDK(e)) : [],
      supply: Array.isArray(object?.supply) ? object.supply.map((e: any) => Coin.fromSDK(e)) : [],
      denomMetadata: Array.isArray(object?.denom_metadata) ? object.denom_metadata.map((e: any) => Metadata.fromSDK(e)) : []
    };
  },
  toSDK(message: GenesisState): GenesisStateSDKType {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toSDK(message.params) : undefined);
    if (message.balances) {
      obj.balances = message.balances.map(e => e ? Balance.toSDK(e) : undefined);
    } else {
      obj.balances = [];
    }
    if (message.supply) {
      obj.supply = message.supply.map(e => e ? Coin.toSDK(e) : undefined);
    } else {
      obj.supply = [];
    }
    if (message.denomMetadata) {
      obj.denom_metadata = message.denomMetadata.map(e => e ? Metadata.toSDK(e) : undefined);
    } else {
      obj.denom_metadata = [];
    }
    return obj;
  },
  fromAmino(object: GenesisStateAmino): GenesisState {
    return {
      params: object?.params ? Params.fromAmino(object.params) : undefined,
      balances: Array.isArray(object?.balances) ? object.balances.map((e: any) => Balance.fromAmino(e)) : [],
      supply: Array.isArray(object?.supply) ? object.supply.map((e: any) => Coin.fromAmino(e)) : [],
      denomMetadata: Array.isArray(object?.denom_metadata) ? object.denom_metadata.map((e: any) => Metadata.fromAmino(e)) : []
    };
  },
  toAmino(message: GenesisState): GenesisStateAmino {
    const obj: any = {};
    obj.params = message.params ? Params.toAmino(message.params) : undefined;
    if (message.balances) {
      obj.balances = message.balances.map(e => e ? Balance.toAmino(e) : undefined);
    } else {
      obj.balances = [];
    }
    if (message.supply) {
      obj.supply = message.supply.map(e => e ? Coin.toAmino(e) : undefined);
    } else {
      obj.supply = [];
    }
    if (message.denomMetadata) {
      obj.denom_metadata = message.denomMetadata.map(e => e ? Metadata.toAmino(e) : undefined);
    } else {
      obj.denom_metadata = [];
    }
    return obj;
  },
  fromAminoMsg(object: GenesisStateAminoMsg): GenesisState {
    return GenesisState.fromAmino(object.value);
  },
  toAminoMsg(message: GenesisState): GenesisStateAminoMsg {
    return {
      type: "cosmos-sdk/GenesisState",
      value: GenesisState.toAmino(message)
    };
  },
  fromProtoMsg(message: GenesisStateProtoMsg): GenesisState {
    return GenesisState.decode(message.value);
  },
  toProto(message: GenesisState): Uint8Array {
    return GenesisState.encode(message).finish();
  },
  toProtoMsg(message: GenesisState): GenesisStateProtoMsg {
    return {
      typeUrl: "/cosmos.bank.v1beta1.GenesisState",
      value: GenesisState.encode(message).finish()
    };
  }
};
function createBaseBalance(): Balance {
  return {
    address: "",
    coins: []
  };
}
export const Balance = {
  typeUrl: "/cosmos.bank.v1beta1.Balance",
  aminoType: "cosmos-sdk/Balance",
  encode(message: Balance, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    for (const v of message.coins) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },
  decode(input: _m0.Reader | Uint8Array, length?: number): Balance {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBalance();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.coins.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },
  fromJSON(object: any): Balance {
    const obj = createBaseBalance();
    if (isSet(object.address)) obj.address = String(object.address);
    if (Array.isArray(object?.coins)) obj.coins = object.coins.map((e: any) => Coin.fromJSON(e));
    return obj;
  },
  toJSON(message: Balance): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    if (message.coins) {
      obj.coins = message.coins.map(e => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.coins = [];
    }
    return obj;
  },
  fromPartial(object: DeepPartial<Balance>): Balance {
    const message = createBaseBalance();
    message.address = object.address ?? "";
    message.coins = object.coins?.map(e => Coin.fromPartial(e)) || [];
    return message;
  },
  fromSDK(object: BalanceSDKType): Balance {
    return {
      address: object?.address,
      coins: Array.isArray(object?.coins) ? object.coins.map((e: any) => Coin.fromSDK(e)) : []
    };
  },
  toSDK(message: Balance): BalanceSDKType {
    const obj: any = {};
    obj.address = message.address;
    if (message.coins) {
      obj.coins = message.coins.map(e => e ? Coin.toSDK(e) : undefined);
    } else {
      obj.coins = [];
    }
    return obj;
  },
  fromAmino(object: BalanceAmino): Balance {
    return {
      address: object.address,
      coins: Array.isArray(object?.coins) ? object.coins.map((e: any) => Coin.fromAmino(e)) : []
    };
  },
  toAmino(message: Balance): BalanceAmino {
    const obj: any = {};
    obj.address = message.address;
    if (message.coins) {
      obj.coins = message.coins.map(e => e ? Coin.toAmino(e) : undefined);
    } else {
      obj.coins = [];
    }
    return obj;
  },
  fromAminoMsg(object: BalanceAminoMsg): Balance {
    return Balance.fromAmino(object.value);
  },
  toAminoMsg(message: Balance): BalanceAminoMsg {
    return {
      type: "cosmos-sdk/Balance",
      value: Balance.toAmino(message)
    };
  },
  fromProtoMsg(message: BalanceProtoMsg): Balance {
    return Balance.decode(message.value);
  },
  toProto(message: Balance): Uint8Array {
    return Balance.encode(message).finish();
  },
  toProtoMsg(message: Balance): BalanceProtoMsg {
    return {
      typeUrl: "/cosmos.bank.v1beta1.Balance",
      value: Balance.encode(message).finish()
    };
  }
};