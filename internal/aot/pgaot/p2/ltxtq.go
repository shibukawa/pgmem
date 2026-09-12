package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltxtq_extract_exec(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v22 = F_array_iterator(m, v12, int32(5654), v18, v9+int32(12))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v26 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v30 != v18 {
								F_pfree(m, v18)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v34 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
									v56 = int32(0)
									m.G0 = v9 + int32(16)
									return v56
								}
							} else {
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v56 = int32(0)
								m.G0 = v9 + int32(16)
								return v56
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v30 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v56 = int32(0)
								m.G0 = v9 + int32(16)
								return v56
							}
						} else {
							v34 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
							v56 = int32(0)
							m.G0 = v9 + int32(16)
							return v56
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v40 = F_palloc0(m, int32(base.Ui32(v37)>>(uint(int32(2))%32)))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
						if v45 != 0 {
							v46 = F__emscripten_memcpy_bulkmem(m, v40, v42, v45)
							mBase = m.M
						} else {
						}
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v48 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v18 == v52 {
									v56 = v40
									m.G0 = v9 + int32(16)
									return v56
								} else {
									F_pfree(m, v18)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = v40
										m.G0 = v9 + int32(16)
										return v56
									}
								}
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v18 == v52 {
								v56 = v40
								m.G0 = v9 + int32(16)
								return v56
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v40
									m.G0 = v9 + int32(16)
									return v56
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ltxtq_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pq_getmsgint(m, v8, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
				F_errmsg_internal(m, int32(471613), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495931), int32(437), int32(36389))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v38 = F_pq_getmsgtext(m, v8, v33-v34, v6+int32(12))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v41 = F_queryin(m, v38, int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v38)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(16)
						return v41
					}
				}
			}
		}
	}
}
