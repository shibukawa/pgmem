package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_series_step_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v15 == int32(3) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(251796), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(515463), int32(1558), int32(583610))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
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
				v21 = v18
				v22 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = int32(4554128)
					v27 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v29
					v32 = F_palloc(m, int32(12))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v21
						*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v12
						*(*int32)(unsafe.Add(mBase, uint32(v32))) = v13
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v32
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
						if int32(0) < v50 {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							if v49 <= v53 {
								v59 = v49 + v50
								*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
								if base.B2i32(v50 < int32(0)) != base.B2i32(v59 < v49) {
									*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(0)
								} else {
								}
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
								*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
								return v49
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
									v96 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
									return int32(0)
								}
							}
						} else {
							if int32(0) <= v50 {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
									v96 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
									return int32(0)
								}
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
								if v49 < v57 {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
										v96 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
										return int32(0)
									}
								} else {
									v59 = v49 + v50
									*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
									if base.B2i32(v50 < int32(0)) != base.B2i32(v59 < v49) {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(0)
									} else {
									}
									v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
									*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
									return v49
								}
							}
						}
					}
				}
			}
		} else {
			v21 = int32(1)
			v22 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = int32(4554128)
				v27 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v29
				v32 = F_palloc(m, int32(12))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v21
					*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v32))) = v13
					*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v32
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v27
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
					if int32(0) < v50 {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
						if v49 <= v53 {
							v59 = v49 + v50
							*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
							if base.B2i32(v50 < int32(0)) != base.B2i32(v59 < v49) {
								*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(0)
							} else {
							}
							v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
							*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
							return v49
						} else {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
								v96 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
								return int32(0)
							}
						}
					} else {
						if int32(0) <= v50 {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
								v96 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
								return int32(0)
							}
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							if v49 < v57 {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
									v96 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
									return int32(0)
								}
							} else {
								v59 = v49 + v50
								*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
								if base.B2i32(v50 < int32(0)) != base.B2i32(v59 < v49) {
									*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(0)
								} else {
								}
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
								*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
								return v49
							}
						}
					}
				}
			}
		}
	} else {
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
		if int32(0) < v50 {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
			if v49 <= v53 {
				v59 = v49 + v50
				*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
				if base.B2i32(v50 < int32(0)) != base.B2i32(v59 < v49) {
					*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(0)
				} else {
				}
				v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
				*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
				return v49
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
					v96 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
					return int32(0)
				}
			}
		} else {
			if int32(0) <= v50 {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
					v96 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
					return int32(0)
				}
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
				if v49 < v57 {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
						v96 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
						return int32(0)
					}
				} else {
					v59 = v49 + v50
					*(*int32)(unsafe.Add(mBase, uint32(v48))) = v59
					if base.B2i32(v50 < int32(0)) != base.B2i32(v59 < v49) {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = int32(0)
					} else {
					}
					v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
					*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
					return v49
				}
			}
		}
	}
}
