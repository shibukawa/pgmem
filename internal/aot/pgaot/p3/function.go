package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FunctionCall5Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v8 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+104)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = l2
	v28 = int32(5)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+26)) = uint16(v28)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v11)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = m.T0[v38].(func(*base.Module, int32) int64)(m, v11+int32(8))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		return int64(0)
	} else {
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
		if v43 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int64(0)
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v50
				F_errmsg_internal(m, int32(_a_F_FunctionCall5Coll_0), v11)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall5Coll_1), int32(1248), int32(_a_F_FunctionCall5Coll_2))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v11 + int32(112)
			return v39
		}
	}
}
func F_FunctionCall7Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64, l7 int64, l8 int64) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v10 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(144)
	m.G0 = v13
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+136)) = uint8(v10)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = l8
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+120)) = uint8(v10)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+112)) = l7
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+104)) = uint8(v10)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+88)) = uint8(v10)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+72)) = uint8(v10)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+56)) = uint8(v10)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+40)) = uint8(v10)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = l2
	v36 = int32(7)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+26)) = uint16(v36)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l0
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = m.T0[v46].(func(*base.Module, int32) int64)(m, v13+int32(8))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return int64(0)
	} else {
		v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
		if v51 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int64(0)
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v58
				F_errmsg_internal(m, int32(_a_F_FunctionCall7Coll_0), v13)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall7Coll_1), int32(1314), int32(_a_F_FunctionCall7Coll_2))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v13 + int32(144)
			return v47
		}
	}
}
func F_coerce_function_result_tuple(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = base.I32_wrap_i64(v11)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 != int32(1) {
		v65 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return
		} else {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v65
			v69 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v69
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+20)) = uint16(v69)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(base.Ui32(v67) >> (uint(int32(2)) % 32))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			v80 = F_lookup_rowtype_tupdesc(m, v78, v79)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				v83 = F_convert_tuples_by_position(m, v80, l1, int32(_a_F_coerce_function_result_tuple_0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					if v83 != 0 {
						v87 = F_execute_attr_map_tuple(m, v9+int32(12), v83)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							v91 = v87
							v92 = F_SPI_returntuple(m, v91, l1)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v92)
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
								if v96 < int32(0) {
									m.G0 = v9 + int32(32)
									return
								} else {
									F_DecrTupleDescRefCount(m, v80)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					} else {
						v91 = v9 + int32(12)
						v92 = F_SPI_returntuple(m, v91, l1)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v92)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
							if v96 < int32(0) {
								m.G0 = v9 + int32(32)
								return
							} else {
								F_DecrTupleDescRefCount(m, v80)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
		if v16&int32(254) != int32(2) {
			v65 = F_pg_detoast_datum(m, v12)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v65
				v69 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v69
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+20)) = uint16(v69)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(base.Ui32(v67) >> (uint(int32(2)) % 32))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
				v80 = F_lookup_rowtype_tupdesc(m, v78, v79)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return
				} else {
					v83 = F_convert_tuples_by_position(m, v80, l1, int32(_a_F_coerce_function_result_tuple_0))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						if v83 != 0 {
							v87 = F_execute_attr_map_tuple(m, v9+int32(12), v83)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								v91 = v87
								v92 = F_SPI_returntuple(m, v91, l1)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v92)
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									if v96 < int32(0) {
										m.G0 = v9 + int32(32)
										return
									} else {
										F_DecrTupleDescRefCount(m, v80)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						} else {
							v91 = v9 + int32(12)
							v92 = F_SPI_returntuple(m, v91, l1)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v92)
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
								if v96 < int32(0) {
									m.G0 = v9 + int32(32)
									return
								} else {
									F_DecrTupleDescRefCount(m, v80)
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v11))+2))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
			if v23 != 0 {
				v26 = v23
				v28 = F_convert_tuples_by_position(m, v26, l1, int32(_a_F_coerce_function_result_tuple_0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					if v28 != 0 {
						v30 = F_expanded_record_get_tuple(m, v22)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							v32 = F_execute_attr_map_tuple(m, v30, v28)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v34 = F_SPI_returntuple(m, v32, l1)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v34)
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
						if v38 == v39 {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
							v62 = F_SPI_datumTransfer(m, v60, int32(-1))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v62
								m.G0 = v9 + int32(32)
								return
							}
						} else {
							if v38 == int32(2249) {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)))
								if v43&int32(64) == int32(0) {
									v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
									v62 = F_SPI_datumTransfer(m, v60, int32(-1))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v62
										m.G0 = v9 + int32(32)
										return
									}
								} else {
									v48 = F_EOH_get_flat_size(m, v22)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										v50 = F_SPI_palloc(m, v48)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											F_EOH_flatten_into(m, v22, v50, v48)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v54
												v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v56
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v50)
												m.G0 = v9 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v48 = F_EOH_get_flat_size(m, v22)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									v50 = F_SPI_palloc(m, v48)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										F_EOH_flatten_into(m, v22, v50, v48)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v54
											v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v56
											*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v50)
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v24 = F_expanded_record_fetch_tupdesc(m, v22)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = v24
					v28 = F_convert_tuples_by_position(m, v26, l1, int32(_a_F_coerce_function_result_tuple_0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						if v28 != 0 {
							v30 = F_expanded_record_get_tuple(m, v22)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								v32 = F_execute_attr_map_tuple(m, v30, v28)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									v34 = F_SPI_returntuple(m, v32, l1)
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v34)
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
							if v38 == v39 {
								v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
								v62 = F_SPI_datumTransfer(m, v60, int32(-1))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v62
									m.G0 = v9 + int32(32)
									return
								}
							} else {
								if v38 == int32(2249) {
									v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)))
									if v43&int32(64) == int32(0) {
										v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
										v62 = F_SPI_datumTransfer(m, v60, int32(-1))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v62
											m.G0 = v9 + int32(32)
											return
										}
									} else {
										v48 = F_EOH_get_flat_size(m, v22)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											v50 = F_SPI_palloc(m, v48)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
											} else {
												F_EOH_flatten_into(m, v22, v50, v48)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v54
													v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v56
													*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v50)
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								} else {
									v48 = F_EOH_get_flat_size(m, v22)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										v50 = F_SPI_palloc(m, v48)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											F_EOH_flatten_into(m, v22, v50, v48)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v54
												v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v56
												*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = base.I64_extend_i32_u(v50)
												m.G0 = v9 + int32(32)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_has_function_privilege_id_name(m *base.Module, l0 int32) int64 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14305(m, l0, int32(_a_F_has_function_privilege_id_name_0), int32(1255), int32(_a_F_has_function_privilege_id_name_1), int32(3589), int32(_a_F_has_function_privilege_id_name_2), int32(52461700), int32(1365))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_has_function_privilege_name_id(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14304(m, l0, int32(_a_F_has_function_privilege_name_id_0), int32(1255))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
