package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FunctionCall5Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
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
	var v39 int32
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
	v11 = v9 + int32(-64)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+52)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l2
	v28 = int32(5)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+22)) = uint16(v28)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = m.T0[v38].(func(*base.Module, int32) int32)(m, v9+int32(-60))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		return int32(0)
	} else {
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
		if v43 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v50
				F_errmsg_internal(m, int32(523043), v11)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(486266), int32(1246), int32(298903))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v11 - int32(-64)
			return v39
		}
	}
}
func F_FunctionCall7Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
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
	var v47 int32
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
	v13 = v11 - int32(80)
	m.G0 = v13
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = l8
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+68)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l7
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+60)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+52)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+44)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+36)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+28)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = l2
	v36 = int32(7)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+22)) = uint16(v36)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)) = uint8(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l0
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = m.T0[v46].(func(*base.Module, int32) int32)(m, v13+int32(4))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return int32(0)
	} else {
		v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)))
		if v51 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v58
				F_errmsg_internal(m, int32(523043), v13)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(486266), int32(1312), int32(298861))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v13 + int32(80)
			return v47
		}
	}
}
func F_coerce_function_result_tuple(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 != int32(1) {
		v60 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v60
			v64 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v64
			*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)) = uint16(v64)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(base.Ui32(v62) >> (uint(int32(2)) % 32))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v75 = F_lookup_rowtype_tupdesc(m, v73, v74)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return
			} else {
				v78 = F_convert_tuples_by_position(m, v75, l1, int32(363996))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					if v78 != 0 {
						v82 = F_execute_attr_map_tuple(m, v8+int32(12), v78)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							v86 = v82
							v87 = F_SPI_returntuple(m, v86, l1)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v87
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
								if v90 < int32(0) {
									m.G0 = v8 + int32(32)
									return
								} else {
									F_DecrTupleDescRefCount(m, v75)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										m.G0 = v8 + int32(32)
										return
									}
								}
							}
						}
					} else {
						v86 = v8 + int32(12)
						v87 = F_SPI_returntuple(m, v86, l1)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v87
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
							if v90 < int32(0) {
								m.G0 = v8 + int32(32)
								return
							} else {
								F_DecrTupleDescRefCount(m, v75)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
		if v14&int32(254) != int32(2) {
			v60 = F_pg_detoast_datum(m, v10)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v60
				v64 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v64
				*(*uint16)(unsafe.Add(mBase, uint32(v8)+20)) = uint16(v64)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(base.Ui32(v62) >> (uint(int32(2)) % 32))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
				v75 = F_lookup_rowtype_tupdesc(m, v73, v74)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					v78 = F_convert_tuples_by_position(m, v75, l1, int32(363996))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						if v78 != 0 {
							v82 = F_execute_attr_map_tuple(m, v8+int32(12), v78)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								v86 = v82
								v87 = F_SPI_returntuple(m, v86, l1)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v87
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
									if v90 < int32(0) {
										m.G0 = v8 + int32(32)
										return
									} else {
										F_DecrTupleDescRefCount(m, v75)
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return
										} else {
											m.G0 = v8 + int32(32)
											return
										}
									}
								}
							}
						} else {
							v86 = v8 + int32(12)
							v87 = F_SPI_returntuple(m, v86, l1)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v87
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
								if v90 < int32(0) {
									m.G0 = v8 + int32(32)
									return
								} else {
									F_DecrTupleDescRefCount(m, v75)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return
									} else {
										m.G0 = v8 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+2))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
			if v20 != 0 {
				v23 = v20
				v25 = F_convert_tuples_by_position(m, v23, l1, int32(363996))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					if v25 != 0 {
						v27 = F_expanded_record_get_tuple(m, v19)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v29 = F_execute_attr_map_tuple(m, v27, v25)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								v31 = F_SPI_returntuple(m, v29, l1)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
						if v34 == v35 {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v57 = F_SPI_datumTransfer(m, v55, int32(-1))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v57
								m.G0 = v8 + int32(32)
								return
							}
						} else {
							if v34 == int32(2249) {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)))
								if v39&int32(64) == int32(0) {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v57 = F_SPI_datumTransfer(m, v55, int32(-1))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v57
										m.G0 = v8 + int32(32)
										return
									}
								} else {
									v44 = F_EOH_get_flat_size(m, v19)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										v46 = F_SPI_palloc(m, v44)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											F_EOH_flatten_into(m, v19, v46, v44)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v50
												v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v52
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v46
												m.G0 = v8 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v44 = F_EOH_get_flat_size(m, v19)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v46 = F_SPI_palloc(m, v44)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										F_EOH_flatten_into(m, v19, v46, v44)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v50
											v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v52
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v46
											m.G0 = v8 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v21 = F_expanded_record_fetch_tupdesc(m, v19)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = v21
					v25 = F_convert_tuples_by_position(m, v23, l1, int32(363996))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						if v25 != 0 {
							v27 = F_expanded_record_get_tuple(m, v19)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								v29 = F_execute_attr_map_tuple(m, v27, v25)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return
								} else {
									v31 = F_SPI_returntuple(m, v29, l1)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
										m.G0 = v8 + int32(32)
										return
									}
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
							if v34 == v35 {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v57 = F_SPI_datumTransfer(m, v55, int32(-1))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v57
									m.G0 = v8 + int32(32)
									return
								}
							} else {
								if v34 == int32(2249) {
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+28)))
									if v39&int32(64) == int32(0) {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v57 = F_SPI_datumTransfer(m, v55, int32(-1))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v57
											m.G0 = v8 + int32(32)
											return
										}
									} else {
										v44 = F_EOH_get_flat_size(m, v19)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return
										} else {
											v46 = F_SPI_palloc(m, v44)
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return
											} else {
												F_EOH_flatten_into(m, v19, v46, v44)
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return
												} else {
													v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v50
													v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v52
													*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v46
													m.G0 = v8 + int32(32)
													return
												}
											}
										}
									}
								} else {
									v44 = F_EOH_get_flat_size(m, v19)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										v46 = F_SPI_palloc(m, v44)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											F_EOH_flatten_into(m, v19, v46, v44)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return
											} else {
												v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v50
												v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v52
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v46
												m.G0 = v8 + int32(32)
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
func F_has_function_privilege_id_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v21 = F_text_to_cstring(m, v12)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_DirectFunctionCall1Coll(m, int32(1252), int32(0), v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v21
								F_errmsg(m, int32(70217), v8)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(488580), int32(3565), int32(372089))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v45 = F_convert_any_priv_string(m, v17, int32(1632528))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = F_object_aclcheck(m, int32(1255), v23, v10, v45)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(16)
								return base.B2i32(v47 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
func F_has_function_privilege_name_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v21 = F_get_role_oid_or_public(m, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = F_convert_any_priv_string(m, v14, int32(1632528))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = F_object_aclcheck_ext(m, int32(1255), v11, v21, v24, v9+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v30 == int32(1) {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
						v37 = int32(0)
					} else {
						v37 = base.B2i32(v28 == int32(0))
					}
					m.G0 = v9 + int32(16)
					return v37
				}
			}
		}
	}
}
