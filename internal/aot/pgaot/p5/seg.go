package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_seg_center(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float32
	_ = v3
	var v4 float32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float32)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*float32)(unsafe.Add(mBase, uint32(v2)+4))
	return base.I32_reinterpret_f32(base.F32_mul(base.F32_add(v3, v4), float32(0.5)))
}
func F_seg_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		F_seg_scanner_init(m, v9, v7+int32(12))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v21 = F_seg_yyparse(m, v11, v19, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v21 != 0 {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					F_seg_yyerror(m, v11, v23, v24, int32(70385))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						F_cube_scanner_finish(m, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return v11
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					F_cube_scanner_finish(m, v28)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v11
					}
				}
			}
		}
	}
}
func F_seg_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(6670), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 <= int32(0))
	}
}
func F_seg_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 float32
	_ = v16
	var v17 float32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float32
	_ = v24
	var v26 float32
	_ = v26
	var v27 float32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 float32
	_ = v34
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_palloc(m, int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*float32)(unsafe.Add(mBase, uint32(v9)+4))
		v17 = *(*float32)(unsafe.Add(mBase, uint32(v10)+4))
		v18 = base.F32_gt(v16, v17)
		if v18 != 0 {
			v19 = v9
		} else {
			v19 = v10
		}
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+9)))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v21)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)) = uint8(v20)
		if v18 != 0 {
			v24 = v16
		} else {
			v24 = v17
		}
		*(*float32)(unsafe.Add(mBase, uint32(v12)+4)) = v24
		v26 = *(*float32)(unsafe.Add(mBase, uint32(v9)))
		v27 = *(*float32)(unsafe.Add(mBase, uint32(v10)))
		v28 = base.F32_lt(v26, v27)
		if v28 != 0 {
			v29 = v9
		} else {
			v29 = v10
		}
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+10)))
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v31)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v30)
		if v28 != 0 {
			v34 = v26
		} else {
			v34 = v27
		}
		*(*float32)(unsafe.Add(mBase, uint32(v12))) = v34
		return v12
	}
}
func F_seg_yy_scan_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v4 = int32(0)
	if base.Ui32(l1) < base.Ui32(int32(2)) {
		v77 = v4
		return v77
	} else {
		v9 = l1 - int32(2)
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v9))))
		if v11 != 0 {
			v77 = v4
			return v77
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1-int32(1)))))
			if v15 != 0 {
				v77 = v4
				return v77
			} else {
				v17 = F_palloc(m, int32(48))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v17 == int32(0) {
						F_yy_fatal_error_7(m, int32(713036))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v23 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v23
						*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v9
						*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(4294967296)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v17))) = v23
						F_seg_yyensure_buffer_stack(m, l2)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38<<(uint(int32(2))%32))))
							if v42 == v17 {
								v77 = v17
							} else {
								if v42 != 0 {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
									v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)))
									*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v45)
									v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48<<(uint(int32(2))%32))))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
									*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v53
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v55
									v57 = v48
									v58 = v47
								} else {
									v57 = v38
									v58 = v37
								}
								*(*int32)(unsafe.Add(mBase, uint32(v58+v57<<(uint(int32(2))%32)))) = v17
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v64
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v66
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v69
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+48)) = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)) = uint8(v71)
								v77 = v17
							}
							return v77
						}
					}
				}
			}
		}
	}
}
func F_seg_yypush_buffer_state(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	if l0 != 0 {
		F_seg_yyensure_buffer_stack(m, l1)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32))))
			if v13 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
				*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v15)
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18<<(uint(int32(2))%32))))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v25
				v28 = v18 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v28
				v30 = v17
				v31 = v28
			} else {
				v30 = v8
				v31 = v9
			}
			*(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)))) = l0
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v37
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v39
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v42
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v44)
			return
		}
	} else {
		return
	}
}
func F_seg_yyrestart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32))))
		if v13 != 0 {
			v27 = v13
			v29 = *(*int32)(unsafe.Add(mBase, _consts[87]))
			v30 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v30)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)) = uint8(v30)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v30
			v40 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v40
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32))))
			if v50 == v27 {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v52
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v54
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v59)
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
				v70 = v40
				v72 = int32(40)
			} else {
				v63 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v63
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v63
				v70 = int32(0)
				v72 = int32(36)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v72+v27))) = v70
			*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[87])) = v29
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v79
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v81
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v81
			v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v86)
			return
		} else {
			F_seg_yyensure_buffer_stack(m, l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v19 = F_seg_yy_create_buffer(m, v17, int32(16384), l1)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)))) = v19
					v27 = v19
					v29 = *(*int32)(unsafe.Add(mBase, _consts[87]))
					v30 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v30
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v30)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)) = uint8(v30)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v30
					v40 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v40
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32))))
					if v50 == v27 {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v52
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v59)
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
						v70 = v40
						v72 = int32(40)
					} else {
						v63 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v63
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v63
						v70 = int32(0)
						v72 = int32(36)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v72+v27))) = v70
					*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[87])) = v29
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v79
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v81
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v81
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84
					v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v86)
					return
				}
			}
		}
	} else {
		F_seg_yyensure_buffer_stack(m, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v19 = F_seg_yy_create_buffer(m, v17, int32(16384), l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)))) = v19
				v27 = v19
				v29 = *(*int32)(unsafe.Add(mBase, _consts[87]))
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v30
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v30)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)) = uint8(v30)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v30
				v40 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v40
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32))))
				if v50 == v27 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v52
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v59)
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
					v70 = v40
					v72 = int32(40)
				} else {
					v63 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v63
					v70 = int32(0)
					v72 = int32(36)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v72+v27))) = v70
				*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(0)
				*(*int32)(unsafe.Add(mBase, _consts[87])) = v29
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v79
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v81
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v81
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84
				v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v86)
				return
			}
		}
	}
}
func F_seg_yyset_column(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4+v5<<(uint(int32(2))%32))))
		if v9 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l0
			return
		} else {
			F_yy_fatal_error_7(m, int32(236278))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		F_yy_fatal_error_7(m, int32(236278))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
