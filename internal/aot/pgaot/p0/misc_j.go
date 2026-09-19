package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_jaccard_distance(m *base.Module, l0 int32) int32 {
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if v16 != v17 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v27
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v26
						F_errmsg(m, int32(_a_F_jaccard_distance_0), v6)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_jaccard_distance_1), int32(39), int32(_a_F_jaccard_distance_2))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
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
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v41 = int32(8)
				v47 = int64(0)
				v51 = *(*int32)(unsafe.Add(mBase, _c_F_jaccard_distance[0]))
				v52 = m.T0[v51].(func(*base.Module, int32, int32, int32, int64, int64, int64) float64)(m, int32(base.Ui32(v38)>>(uint(int32(2))%32))-v41, v9+v41, v14+v41, v47, v47, v47)
				mBase = m.M
				v53 = F_Float8GetDatum(m, v52)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v53
				}
			}
		}
	}
}
func F_johab_to_utf8(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn13870(m, l0, int32(40), v3, v3, v3, int32(_a_F_johab_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_jsonpath_out(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v16 = v7 + int32(32)
		F_initStringInfo(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_enlargeStringInfo(m, v16, int32(base.Ui32(v14)>>(uint(int32(2))%32)))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if int32(0) <= v23 {
					F_appendStringInfoString(m, v16, int32(_a_F_jsonpath_out_0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v30 = v7 + int32(4)
						F_jspInitByBuffer(m, v30, v10+int32(8), int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_printJsonPathItem(m, v7+int32(32), v30, int32(0), int32(1))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
								m.G0 = v7 + int32(48)
								return v42
							}
						}
					}
				} else {
					v30 = v7 + int32(4)
					F_jspInitByBuffer(m, v30, v10+int32(8), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_printJsonPathItem(m, v7+int32(32), v30, int32(0), int32(1))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
							m.G0 = v7 + int32(48)
							return v42
						}
					}
				}
			}
		}
	}
}
func F_jsonpath_send(m *base.Module, l0 int32) int32 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v6 + int32(4)
		F_initStringInfo(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			F_enlargeStringInfo(m, v14, int32(base.Ui32(v17)>>(uint(int32(2))%32)))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				if int32(0) <= v22 {
					F_appendStringInfoString(m, v14, int32(_a_F_jsonpath_send_0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v29 = v6 + int32(20)
						F_jspInitByBuffer(m, v29, v9+int32(8), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_printJsonPathItem(m, v6+int32(4), v29, int32(0), int32(1))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								F_pq_begintypsend(m, v29)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_enlargeStringInfo(m, v29, int32(1))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
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
											return int32(0)
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
											F_pfree(m, v58)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v62))) = v63 << (uint(int32(2)) % 32)
												m.G0 = v6 + int32(48)
												return v62
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
						return int32(0)
					} else {
						F_printJsonPathItem(m, v6+int32(4), v29, int32(0), int32(1))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_pq_begintypsend(m, v29)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_enlargeStringInfo(m, v29, int32(1))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
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
										return int32(0)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
										F_pfree(m, v58)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v62))) = v63 << (uint(int32(2)) % 32)
											m.G0 = v6 + int32(48)
											return v62
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
