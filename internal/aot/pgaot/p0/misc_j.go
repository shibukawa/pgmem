package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_jaccard_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if v17 != v18 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v28
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
						F_errmsg(m, int32(_a_F_jaccard_distance_0), v7)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_jaccard_distance_1), int32(39), int32(_a_F_jaccard_distance_2))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v42 = int32(8)
				v48 = int64(0)
				v52 = *(*int32)(unsafe.Add(mBase, _c_F_jaccard_distance[0]))
				v53 = m.T0[v52].(func(*base.Module, int32, int32, int32, int64, int64, int64) float64)(m, int32(base.Ui32(v39)>>(uint(int32(2))%32))-v42, v10+v42, v15+v42, v48, v48, v48)
				mBase = m.M
				m.G0 = v7 + int32(16)
				return base.I64_reinterpret_f64(v53)
			}
		}
	}
}
func F_jbv_to_infunc_datum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_jbv_to_infunc_datum[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v14
	v17 = *(*int64)(unsafe.Add(mBase, _c_F_jbv_to_infunc_datum[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = F_palloc0(m, v19+int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v26 != 0 {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			base.MemoryCopy(m, v22, v27, v26)
		} else {
		}
		v32 = F_DirectInputFunctionCallSafe(m, l1, v22, int32(-1), v11+int32(32), l4)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			if v32 == int32(0) {
				v37 = v11 + int32(16)
				F_initStringInfo(m, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = l3
					F_appendStringInfo(m, v37, int32(_a_F_jbv_to_infunc_datum_0), v11)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v45))) = int32(19)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v48)+44)) = v49
						F_ThrowErrorData(m, v48)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
							F_pfree(m, v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v22)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									m.G0 = v11 + int32(48)
									return v32
								}
							}
						}
					}
				}
			} else {
				F_pfree(m, v22)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					m.G0 = v11 + int32(48)
					return v32
				}
			}
		}
	}
}
func F_johab_to_utf8(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14231(m, l0, int32(40), v3, v3, v3, int32(_a_F_johab_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_jsonpath_out(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v17 = v8 + int32(32)
		F_initStringInfo(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			F_enlargeStringInfo(m, v17, int32(base.Ui32(v15)>>(uint(int32(2))%32)))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				if int32(0) <= v24 {
					F_appendStringInfoString(m, v17, int32(_a_F_jsonpath_out_0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v31 = v8 + int32(4)
						F_jspInitByBuffer(m, v31, v11+int32(8), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							F_printJsonPathItem(m, v8+int32(32), v31, int32(0), int32(1))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								v43 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+32)))
								m.G0 = v8 + int32(48)
								return v43
							}
						}
					}
				} else {
					v31 = v8 + int32(4)
					F_jspInitByBuffer(m, v31, v11+int32(8), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						F_printJsonPathItem(m, v8+int32(32), v31, int32(0), int32(1))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							v43 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+32)))
							m.G0 = v8 + int32(48)
							return v43
						}
					}
				}
			}
		}
	}
}
func F_jsonpath_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v14 = v6 + int32(4)
		F_initStringInfo(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			F_enlargeStringInfo(m, v14, int32(base.Ui32(v17)>>(uint(int32(2))%32)))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				if int32(0) <= v22 {
					F_appendStringInfoString(m, v14, int32(_a_F_jsonpath_send_0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v29 = v6 + int32(20)
						F_jspInitByBuffer(m, v29, v9+int32(8), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							F_printJsonPathItem(m, v6+int32(4), v29, int32(0), int32(1))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								F_pq_begintypsend(m, v29)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									F_enlargeStringInfo(m, v29, int32(1))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int64(0)
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
										v49 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v46+v47))) = uint8(v49)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v46 + v49
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
										F_pq_sendtext(m, v29, v54, v55)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int64(0)
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
											F_pfree(m, v58)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int64(0)
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v62))) = v63 << (uint(int32(2)) % 32)
												m.G0 = v6 + int32(48)
												return base.I64_extend_i32_u(v62)
											}
										}
									}
								}
							}
						}
					}
				} else {
					v29 = v6 + int32(20)
					F_jspInitByBuffer(m, v29, v9+int32(8), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						F_printJsonPathItem(m, v6+int32(4), v29, int32(0), int32(1))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							F_pq_begintypsend(m, v29)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								F_enlargeStringInfo(m, v29, int32(1))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int64(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
									v49 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v46+v47))) = uint8(v49)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v46 + v49
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
									F_pq_sendtext(m, v29, v54, v55)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int64(0)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
										F_pfree(m, v58)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int64(0)
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v62))) = v63 << (uint(int32(2)) % 32)
											m.G0 = v6 + int32(48)
											return base.I64_extend_i32_u(v62)
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
