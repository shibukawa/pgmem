package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gistPopItupFromNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v13 <= int32(0) {
		m.G0 = v11 + int32(16)
		return base.B2i32(int32(0) < v13)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		if v16 == int32(0) {
			F_gistLoadNodeBuffer(m, l0, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v24 = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v26 = v24 + v25
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+14)))
				v29 = v27 & int32(_a_F_gistPopItupFromNodeBuffer_0)
				v30 = F_palloc(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
					if v29 != 0 {
						base.MemoryCopy(m, v30, v26+int32(8), v29)
					} else {
					}
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v36 + (v29+int32(7))&int32(_a_F_gistPopItupFromNodeBuffer_1)
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					if v44 != int32(_a_F_gistPopItupFromNodeBuffer_2) {
						m.G0 = v11 + int32(16)
						return base.B2i32(int32(0) < v13)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v47 - int32(1)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v51 != int32(-1) {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v56 = F_BufFileSeekBlock(m, v54, base.I64_extend_i32_s(v51))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51
										F_errmsg_internal(m, int32(_a_F_gistPopItupFromNodeBuffer_3), v11)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_gistPopItupFromNodeBuffer_4), int32(749), int32(_a_F_gistPopItupFromNodeBuffer_5))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									F_BufFileReadExact(m, v54, v43, int32(_a_F_gistPopItupFromNodeBuffer_6))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v61 < v62 {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v75 = v61
											v76 = v64
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v75 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v76+v75<<(uint(int32(2))%32)))) = v51
											m.G0 = v11 + int32(16)
											return base.B2i32(int32(0) < v13)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62 << (uint(int32(1)) % 32)
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											v71 = F_repalloc(m, v68, v62<<(uint(int32(3))%32))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v71
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												v75 = v74
												v76 = v71
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v75 + int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v76+v75<<(uint(int32(2))%32)))) = v51
												m.G0 = v11 + int32(16)
												return base.B2i32(int32(0) < v13)
											}
										}
									}
								}
							}
						} else {
							F_pfree(m, v43)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
								m.G0 = v11 + int32(16)
								return base.B2i32(int32(0) < v13)
							}
						}
					}
				}
			}
		} else {
			v24 = v16
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v26 = v24 + v25
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+14)))
			v29 = v27 & int32(_a_F_gistPopItupFromNodeBuffer_0)
			v30 = F_palloc(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
				if v29 != 0 {
					base.MemoryCopy(m, v30, v26+int32(8), v29)
				} else {
				}
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v36 + (v29+int32(7))&int32(_a_F_gistPopItupFromNodeBuffer_1)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				if v44 != int32(_a_F_gistPopItupFromNodeBuffer_2) {
					m.G0 = v11 + int32(16)
					return base.B2i32(int32(0) < v13)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v47 - int32(1)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					if v51 != int32(-1) {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v56 = F_BufFileSeekBlock(m, v54, base.I64_extend_i32_s(v51))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							if v56 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51
									F_errmsg_internal(m, int32(_a_F_gistPopItupFromNodeBuffer_3), v11)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_gistPopItupFromNodeBuffer_4), int32(749), int32(_a_F_gistPopItupFromNodeBuffer_5))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								F_BufFileReadExact(m, v54, v43, int32(_a_F_gistPopItupFromNodeBuffer_6))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v61 < v62 {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v75 = v61
										v76 = v64
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v75 + int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v76+v75<<(uint(int32(2))%32)))) = v51
										m.G0 = v11 + int32(16)
										return base.B2i32(int32(0) < v13)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v62 << (uint(int32(1)) % 32)
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v71 = F_repalloc(m, v68, v62<<(uint(int32(3))%32))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v71
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											v75 = v74
											v76 = v71
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v75 + int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v76+v75<<(uint(int32(2))%32)))) = v51
											m.G0 = v11 + int32(16)
											return base.B2i32(int32(0) < v13)
										}
									}
								}
							}
						}
					} else {
						F_pfree(m, v43)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
							m.G0 = v11 + int32(16)
							return base.B2i32(int32(0) < v13)
						}
					}
				}
			}
		}
	}
}
func F_gistProcessEmptyingQueue(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = v16
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v13 + int32(16)
	return
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = F_list_delete_first(m, v18)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v29
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+16)) = uint8(v32)
	v35 = m.G0
	v37 = v35 - int32(16)
	m.G0 = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	if v32 < v39 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v125 = F_gistPopItupFromNodeBuffer(m, v15, v28, v13+int32(12))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L31
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L27
	}
L10:
	;
	v43 = v39
	v48 = v32
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = int32(0)
	m.G0 = v37 + int32(16)
	goto L8
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v48<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v57 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if int32(0) < v58 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v89 = v43
	goto L17
L17:
	;
	v93 = v48 + int32(1)
	if v93 < v89 {
		v43 = v89
		v48 = v93
		goto L13
	} else {
		goto L26
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v77 = F_BufFileSeekBlock(m, v75, base.I64_extend_i32_s(v73))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L22
	}
L19:
	;
	v62 = v58 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v62<<(uint(int32(2))%32))))
	v73 = v68
	goto L18
L20:
	;
	goto L21
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v69 + int32(1)
	v73 = v69
	goto L18
L22:
	;
	if v77 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_BufFileWrite(m, v75, v74, int32(_a_F_gistProcessEmptyingQueue_0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	F_pfree(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v89 = v88
	goto L17
L26:
	;
	goto L14
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v73
	F_errmsg_internal(m, int32(_a_F_gistProcessEmptyingQueue_1), v37)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_gistProcessEmptyingQueue_2), int32(757), int32(_a_F_gistProcessEmptyingQueue_3))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v162 != 0 {
		v18 = v162
		goto L4
	} else {
		goto L40
	}
L31:
	;
	if v125 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	goto L33
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v142 = F_gistProcessItup(m, l0, v139, v140, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L35
	}
L34:
	;
	goto L30
L35:
	;
	if v142 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	F_MemoryContextReset(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v150 = F_gistPopItupFromNodeBuffer(m, v15, v28, v13+int32(12))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	if v150 != 0 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	goto L5
}
func F_gistPushItupToNodeBuffer(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(_a_F_gistPushItupToNodeBuffer_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_gistPushItupToNodeBuffer[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _c_F_gistPushItupToNodeBuffer[0])) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v54 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v20 = F_MemoryContextAllocZero(m, v16, int32(_a_F_gistPushItupToNodeBuffer_6))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(35154307317759)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v20
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v28 < v29 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v42<<(uint(int32(2))%32)))) = l1
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v48 + int32(1)
	goto L1
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v42 = v28
	v43 = v31
	goto L6
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v29 << (uint(int32(1)) % 32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v38 = F_repalloc(m, v35, v29<<(uint(int32(3))%32))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v42 = v41
	v43 = v38
	goto L6
L11:
	;
	F_gistLoadNodeBuffer(m, l0, l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L14
	}
L12:
	;
	v60 = v54
	goto L13
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v64 = v62 & int32(_a_F_gistPushItupToNodeBuffer_1)
	if base.Ui32(v61) < base.Ui32((v64+int32(7))&int32(_a_F_gistPushItupToNodeBuffer_2)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v60 = v59
	goto L13
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L33
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v70 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v108 = v60
	v109 = v64
	v110 = v61
	goto L18
L18:
	;
	v115 = v110 - (v109+int32(7))&int32(_a_F_gistPushItupToNodeBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v115
	if v109 != 0 {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v89 = F_BufFileSeekBlock(m, v87, base.I64_extend_i32_s(v85))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L23
	}
L20:
	;
	v74 = v70 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v74<<(uint(int32(2))%32))))
	v85 = v80
	goto L19
L21:
	;
	goto L22
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81 + int32(1)
	v85 = v81
	goto L19
L23:
	;
	if v89 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	F_BufFileWrite(m, v87, v86, int32(_a_F_gistPushItupToNodeBuffer_6))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = int32(_a_F_gistPushItupToNodeBuffer_7)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v85
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v99 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v108 = v103
	v109 = v105 & int32(_a_F_gistPushItupToNodeBuffer_1)
	v110 = v104
	goto L18
L26:
	;
	base.MemoryCopy(m, v108+v115+int32(8), l2, v109)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v124 = base.I32_div_s(v122, int32(2))
	if v121 <= v124 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gistPushItupToNodeBuffer[0])) = v14
	m.G0 = v11 + int32(16)
	return
L30:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v126 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v128 = F_lcons(m, l1, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v128
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v131)
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v85
	F_errmsg_internal(m, int32(_a_F_gistPushItupToNodeBuffer_3), v11)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_gistPushItupToNodeBuffer_4), int32(757), int32(_a_F_gistPushItupToNodeBuffer_5))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_bbox_zorder_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	return int32(0)
}
func F_gist_page_items(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v169 int64
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int64
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v433 int64
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v467 int32
	_ = v467
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	v16 = m.G0
	v18 = v16 - int32(400)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = F_superuser(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L110
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L106
	}
L5:
	;
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L102
	}
L9:
	;
	v33 = F_index_open(m, v26, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+84))
	if v36 != int32(783) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v39 = F_verify_gist_page(m, v21)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	m.G0 = v18 + int32(400)
	return int64(0)
L13:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+14)))
	if v41 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_relation_close(m, v33, int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v49 = int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+16)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v51)+12)))
	if v53&v49 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v47)
	goto L12
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)+192))
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+10)))
	v61 = F_CreateTupleDescTruncatedCopy(m, v50, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v63 = v49
	v64 = v50
	goto L20
L20:
	;
	v65 = int32(0)
	v67 = int32(1)
	v68 = int32(2)
	if v63&v67 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v63 = int32(3)
	v64 = v61
	goto L20
L22:
	;
	v78 = int32(7)
	goto L24
L23:
	;
	v78 = v68
	goto L24
L24:
	;
	v80 = F_pg_get_indexdef_worker(m, v26, v65, v65, v67, int32(base.Ui32(v63&v68)>>(uint(v67)%32)), v65, v65, v78, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+16)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v82)+12)))
	if v84&int32(2) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_relation_close(m, v33, int32(1))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L101
	}
L27:
	;
	v89 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+12)))
	if base.Ui32(v102) < base.Ui32(int32(25)) {
		goto L26
	} else {
		goto L34
	}
L30:
	;
	if v89 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	F_errmsg_internal(m, int32(_a_F_gist_page_items_0), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_1), int32(258), int32(_a_F_gist_page_items_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	v106 = v102 + int32(_a_F_gist_page_items_3)
	if v106&int32(_a_F_gist_page_items_4) == int32(0) {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v117 = int32(1)
	v119 = v117
	v127 = v117
	goto L36
L36:
	;
	v136 = v39 + int32(20) + v119<<(uint(int32(2))%32)
	if v136 == int32(0) {
		goto L3
	} else {
		goto L38
	}
L37:
	;
	goto L26
L38:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v142 = v39 + v139&int32(_a_F_gist_page_items_5)
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+6)))
	if int32(0) <= base.I32_extend16_s(v149) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+348)) = uint8(v160)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+344)) = v160
	*(*int64)(unsafe.Add(mBase, uint32(v18)+360)) = base.I64_extend_i32_u(v142)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+352)) = base.I64_extend16_s(base.I64_extend_i32_u(v127))
	v169 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v142)+6)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+368)) = v169 & int64(8191)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v174 = int32(_a_F_gist_page_items_6)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+376)) = base.I64_extend_i32_u(base.B2i32(v173&v174 == v174))
	if v80 == v160 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v153 = int32(8)
	goto L42
L41:
	;
	v153 = int32(16)
	goto L42
L42:
	;
	F_index_deform_tuple_internal(m, v64, v18+int32(80), v18+int32(48), v142+v153, v142+int32(8), int32(base.Ui32(v149)>>(uint(int32(15))%32)))
	mBase = m.M
	goto L39
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+348)) = uint8(v434)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+384)) = v433
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	F_tuplestore_putvalues(m, v437, v438, v18+int32(352), v18+int32(344))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L99
	}
L44:
	;
	v433 = int64(0)
	v434 = int32(1)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v185 = v18 + int32(32)
	F_initStringInfo(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v80
	F_appendStringInfo(m, v185, int32(_a_F_gist_page_items_7), v18)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v192 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v192 < v193 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v196 = v193
	v202 = v192
	goto L52
L50:
	;
	goto L51
L51:
	;
	F_appendStringInfoChar(m, v18+int32(32), int32(41))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L97
	}
L52:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(48)+v202))))
	if v214 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L51
L54:
	;
	v238 = int32(_a_F_gist_page_items_8)
	goto L56
L55:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v64+v196<<(uint(int32(3))%32)+v202*int32(100))+96))
	F_getTypeOutputInfo(m, v222, v18+int32(28), v18+int32(27))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v33)+192))
	v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v241)+10)))
	if v242 == v202 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v18+int32(80)+v202<<(uint(int32(3))%32))))
	v236 = F_OidOutputFunctionCall(m, v229, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v238 = v236
	goto L56
L59:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	v254 = v238
	v255 = v251
	goto L67
L60:
	;
	v248 = int32(_a_F_gist_page_items_9)
	goto L62
L61:
	;
	if v202 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L62:
	;
	F_appendStringInfoString(m, v18+int32(32), v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	v248 = int32(_a_F_gist_page_items_10)
	goto L62
L64:
	;
	goto L59
L65:
	;
	v297 = v238
	goto L77
L66:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v271 <= v272+int32(1) {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	switch v255 {
	case 0:
		goto L69
	default:
		goto L70
	case 9, 10, 11, 12, 13, 32, 34, 40, 41, 44, 92:
		goto L66
	}
L68:
	;
	if v251 != 0 {
		v296 = int32(0)
		goto L65
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	v254 = v254 + int32(1)
	v255 = v267
	goto L67
L71:
	;
	goto L66
L72:
	;
	F_appendStringInfoChar(m, v18+int32(32), int32(34))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v284 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v282+v272))) = uint8(v284)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v287 = int32(1)
	v288 = v286 + v287
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v292 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v290+v288))) = uint8(v292)
	v296 = v287
	goto L65
L75:
	;
	v296 = int32(1)
	goto L65
L76:
	;
	v391 = v202 + int32(1)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v391 < v392 {
		v196 = v392
		v202 = v391
		goto L52
	} else {
		goto L96
	}
L77:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	v313 = base.I32_extend8_s(v312)
	if base.B2i32(v312 == int32(34))|base.B2i32(v312 == int32(92)) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v379 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v377+v324))) = uint8(v379)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v383 = v381 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v387 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v385+v383))) = uint8(v387)
	goto L76
L79:
	;
	goto L78
L80:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v354 <= v355+int32(1) {
		goto L92
	} else {
		goto L93
	}
L81:
	;
	if v312 != 0 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v333 <= v334+int32(1) {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	if v296 == int32(0) {
		goto L76
	} else {
		goto L85
	}
L85:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v324+int32(1) < v323 {
		goto L79
	} else {
		goto L86
	}
L86:
	;
	F_appendStringInfoChar(m, v18+int32(32), int32(34))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L76
L88:
	;
	F_appendStringInfoChar(m, v18+int32(32), v313)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v342+v334))) = uint8(v313)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v347 = v345 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v347
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v351 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v349+v347))) = uint8(v351)
	goto L80
L91:
	;
	goto L80
L92:
	;
	F_appendStringInfoChar(m, v18+int32(32), v313)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v363+v355))) = uint8(v313)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v368 = v366 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v368
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v372 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v370+v368))) = uint8(v372)
	goto L94
L94:
	;
	v297 = v297 + int32(1)
	goto L77
L95:
	;
	goto L94
L96:
	;
	goto L53
L97:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v415 = F_cstring_to_text(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v433 = base.I64_extend_i32_u(v415)
	v434 = int32(0)
	goto L43
L99:
	;
	v446 = v127 + int32(1)
	v448 = v446 & int32(_a_F_gist_page_items_11)
	if base.Ui32(v448) <= base.Ui32(int32(base.Ui32(v106)>>(uint(int32(2))%32))&int32(_a_F_gist_page_items_11)) {
		v119 = v448
		v127 = v446
		goto L36
	} else {
		goto L100
	}
L100:
	;
	goto L37
L101:
	;
	goto L12
L102:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_gist_page_items_12), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_1), int32(214), int32(_a_F_gist_page_items_2))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(_a_F_gist_page_items_13)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v511 + int32(4)
	F_errmsg(m, int32(_a_F_gist_page_items_14), v18+int32(16))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_1), int32(225), int32(_a_F_gist_page_items_2))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errmsg_internal(m, int32(_a_F_gist_page_items_15), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_1), int32(278), int32(_a_F_gist_page_items_2))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_page_items_bytea(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v94 int32
	_ = v94
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = F_superuser(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L36
	}
L4:
	;
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L32
	}
L8:
	;
	v28 = F_verify_gist_page(m, v18)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v15 - int32(-64)
	return int64(0)
L10:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+14)))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
	goto L9
L12:
	;
	goto L13
L13:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v35)+12)))
	if v37&int32(2) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+12)))
	if base.Ui32(v55) < base.Ui32(int32(25)) {
		goto L9
	} else {
		goto L21
	}
L17:
	;
	if v42 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	F_errmsg_internal(m, int32(_a_F_gist_page_items_bytea_0), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_bytea_1), int32(155), int32(_a_F_gist_page_items_bytea_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	v59 = v55 + int32(_a_F_gist_page_items_bytea_3)
	if v59&int32(_a_F_gist_page_items_bytea_4) == int32(0) {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v70 = int32(1)
	v72 = v70
	v75 = v70
	goto L23
L23:
	;
	v86 = v28 + int32(20) + v75<<(uint(int32(2))%32)
	if v86 == int32(0) {
		goto L3
	} else {
		goto L25
	}
L24:
	;
	goto L9
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v92 = v28 + v89&int32(_a_F_gist_page_items_bytea_5)
	v93 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v92)+6)))
	v94 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = base.I64_extend_i32_u(v92)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = base.I64_extend16_s(base.I64_extend_i32_u(v72))
	v104 = v93 & int64(8191)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = v104
	v106 = base.I32_wrap_i64(v104)
	v108 = v106 + int32(4)
	v109 = F_palloc(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v108 << (uint(int32(2)) % 32)
	if v106 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	base.MemoryCopy(m, v109+int32(4), v92, v106)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+48)) = base.I64_extend_i32_u(v109)
	v120 = int32(_a_F_gist_page_items_bytea_6)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = base.I64_extend_i32_u(base.B2i32(v117&v120 == v120))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_tuplestore_putvalues(m, v126, v127, v13+int32(-48), v13+int32(-56))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v135 = v72 + int32(1)
	v137 = v135 & int32(_a_F_gist_page_items_bytea_7)
	if base.Ui32(v137) <= base.Ui32(int32(base.Ui32(v59)>>(uint(int32(2))%32))&int32(_a_F_gist_page_items_bytea_7)) {
		v72 = v135
		v75 = v137
		goto L23
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_gist_page_items_bytea_8), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_bytea_1), int32(144), int32(_a_F_gist_page_items_bytea_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errmsg_internal(m, int32(_a_F_gist_page_items_bytea_9), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_gist_page_items_bytea_1), int32(173), int32(_a_F_gist_page_items_bytea_2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_page_opaque_info(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v7 = m.G0
	v9 = v7 - int32(192)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = F_superuser(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L46
	}
L4:
	;
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v18 = F_verify_gist_page(m, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L8:
	;
	m.G0 = v9 + int32(192)
	return v144
L9:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+14)))
	if v20 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
	v144 = int64(0)
	goto L8
L11:
	;
	goto L12
L12:
	;
	v26 = int32(0)
	v30 = F_get_call_result_type(m, l0, v26, v9+int32(188))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v30 != int32(1) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)))
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+v34)+12)))
	if v36&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_3))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v47 = v9
	v48 = v26
	goto L17
L17:
	;
	if v36&int32(2) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = base.I64_extend_i32_u(v42)
	v47 = v9 | int32(8)
	v48 = int32(1)
	goto L17
L19:
	;
	v52 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_4))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v58 = v48
	goto L21
L21:
	;
	if v36&int32(4) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = base.I64_extend_i32_u(v52)
	v58 = v48 + int32(1)
	goto L21
L23:
	;
	v65 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_5))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v71 = v58
	goto L25
L25:
	;
	if v36&int32(8) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+v58<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v65)
	v71 = v58 + int32(1)
	goto L25
L27:
	;
	v78 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_6))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v84 = v71
	goto L29
L29:
	;
	if v36&int32(16) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+v71<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v78)
	v84 = v71 + int32(1)
	goto L29
L31:
	;
	v91 = F_cstring_to_text(m, int32(_a_F_gist_page_opaque_info_7))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	v97 = v84
	goto L33
L33:
	;
	v99 = v36 & int32(_a_F_gist_page_opaque_info_8)
	if v99 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+v84<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v91)
	v97 = v84 + int32(1)
	goto L33
L35:
	;
	v106 = F_DirectFunctionCall1Coll(m, int32(3109), int32(0), base.I64_extend_i32_u(v99))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v111 = v97
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = int32(0)
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v115 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+144)) = base.I64_rotl(v114, v115)
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)))
	v119 = v18 + v118
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = base.I64_rotl(v120, v115)
	v124 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v119)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+160)) = v124
	v127 = F_construct_array_builtin(m, v9, v111, int32(25))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9+v97<<(uint(int32(3))%32)))) = v106
	v111 = v97 + int32(1)
	goto L37
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+168)) = base.I64_extend_i32_u(v127)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
	v136 = F_heap_form_tuple(m, v131, v9+int32(144), v9+int32(140))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v139 = F_HeapTupleHeaderGetDatum(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v144 = v139
	goto L8
L42:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_gist_page_opaque_info_9), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_gist_page_opaque_info_1), int32(89), int32(_a_F_gist_page_opaque_info_2))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errmsg_internal(m, int32(_a_F_gist_page_opaque_info_0), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_gist_page_opaque_info_1), int32(98), int32(_a_F_gist_page_opaque_info_2))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gist_point_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	if v5 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = int32(115)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(116)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(117)
		v15 = int32(118)
	} else {
		v15 = int32(115)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v15
	return int64(0)
}
func F_gist_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int64
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int64
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int64
	_ = v306
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int64
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int64
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v503 int32
	_ = v503
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v665 int64
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int64
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int64
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int64
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int64
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int64
	_ = v831
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(96)
	m.G0 = v25
	v27 = int32(_a_F_gist_redo_0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_gist_redo[0])) = v33
	v36 = v30 & int32(240)
	switch int32(base.Ui32(v36) >> (uint(int32(4)) % 32)) {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	default:
		goto L1
	case 6:
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L10
	} else {
		goto L192
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L10
	} else {
		goto L189
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L10
	} else {
		goto L186
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gist_redo[0])) = v28
	v938 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[1]))
	F_MemoryContextReset(m, v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L10
	} else {
		goto L185
	}
L5:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v823 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v827 = F_XLogReadBufferForRedo(m, l0, int32(0), v25+int32(92))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L10
	} else {
		goto L160
	}
L6:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379)+18)))
	if v380 != 0 {
		goto L84
	} else {
		goto L85
	}
L7:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[2]))
	if base.Ui32(v365) < base.Ui32(int32(2)) {
		goto L4
	} else {
		goto L82
	}
L8:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v290 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[2]))
	if base.Ui32(int32(2)) <= base.Ui32(v292) {
		goto L65
	} else {
		goto L66
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v44 = F_XLogReadBufferForRedo(m, l0, int32(0), v25+int32(92))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v44 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = int32(0)
	v50 = v25 + int32(76)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+72))
	if v53 < v48 {
		v75 = v48
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+72))
	if v223 <= int32(0) {
		goto L49
	} else {
		goto L50
	}
L15:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	if v79 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	v78 = v75
	goto L15
L17:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(0))+76)))
	if v58 != int32(1) {
		v75 = v48
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v62 = v52 + int32(76)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+43)))
	if v63 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v50 == int32(0) {
		v75 = v48
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v50 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v68 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v68
	v78 = v68
	goto L15
L23:
	;
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v71
	goto L25
L24:
	;
	goto L25
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	v75 = v73
	goto L16
L26:
	;
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	switch v98 {
	case 0:
		v118 = v78
		goto L30
	case 1:
		goto L32
	default:
		goto L31
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+(v79^int32(-1))<<(uint(int32(2))%32))))
	v97 = v89
	goto L26
L28:
	;
	goto L29
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v97 = v91 + v79<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if base.Ui32(v118-v78) < base.Ui32(v120) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	F_PageIndexMultiDelete(m, v97, v78, v98)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L36
	}
L32:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+2)))
	if v99 != int32(1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78))))
	v104 = v78 + int32(2)
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+8)))
	v107 = v105 & int32(_a_F_gist_redo_1)
	v108 = F_PageIndexTupleOverwrite(m, v97, v102, v104, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	if v108 == int32(0) {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v118 = v107 + v104
	goto L30
L36:
	;
	v118 = v78 + v98<<(uint(int32(1))%32)
	goto L30
L37:
	;
	v123 = int32(1)
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+12)))
	if base.Ui32(v124) < base.Ui32(int32(25)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = base.I64_rotl(v40, int64(32))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	F_MarkBufferDirty(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L10
	} else {
		goto L48
	}
L40:
	;
	v133 = v123
	goto L42
L41:
	;
	v133 = int32(base.Ui32(v124+int32(_a_F_gist_redo_2))>>(uint(int32(2))%32)) + v123
	goto L42
L42:
	;
	v135 = v118
	v140 = v133
	goto L43
L43:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+6)))
	v158 = v156 & int32(_a_F_gist_redo_1)
	v162 = F_PageAddItemExtended(m, v97, v135, v158, v140&int32(_a_F_gist_redo_3), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L10
	} else {
		goto L45
	}
L44:
	;
	goto L39
L45:
	;
	if v162 == int32(0) {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v169 = v135 + v158
	if base.Ui32(v169-v78) < base.Ui32(v168) {
		v135 = v169
		v140 = v140 + int32(1)
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	goto L14
L49:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	if v284 == int32(0) {
		goto L4
	} else {
		goto L63
	}
L50:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+128)))
	if v226 != int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v229 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v233 = F_XLogReadBufferForRedo(m, l0, int32(1), v25+int32(76))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	if v233&int32(-3) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if v239 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if v276 == int32(0) {
		goto L49
	} else {
		goto L61
	}
L56:
	;
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+16)))
	v261 = base.I64_rotl(v229, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v258+v257))) = v261
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+16)))
	v264 = v257 + v263
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264)+12)))
	v267 = v265 & int32(_a_F_gist_redo_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v264)+12)) = uint16(v267)
	*(*int64)(unsafe.Add(mBase, uint32(v257))) = v261
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	F_MarkBufferDirty(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L10
	} else {
		goto L60
	}
L57:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v243+(v239^int32(-1))<<(uint(int32(2))%32))))
	v257 = v249
	goto L56
L58:
	;
	goto L59
L59:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v257 = v251 + v239<<(uint(int32(13))%32) + int32(-8192)
	goto L56
L60:
	;
	goto L55
L61:
	;
	F_UnlockReleaseBuffer(m, v276)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	goto L49
L63:
	;
	F_UnlockReleaseBuffer(m, v284)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	goto L4
L65:
	;
	v295 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v295, v25+int32(76), v295, v295)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L10
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v316 = F_XLogReadBufferForRedo(m, l0, int32(0), v25+int32(76))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L10
	} else {
		goto L70
	}
L68:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+6)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = v304
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v25)+76))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v306
	F_ResolveRecoveryConflictWithSnapshot(m, v303, v302, v25+int32(48))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	if v316 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if v322 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	goto L73
L73:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if v359 == int32(0) {
		goto L4
	} else {
		goto L80
	}
L74:
	;
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289)+4)))
	F_PageIndexMultiDelete(m, v340, v289+int32(8), v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L10
	} else {
		goto L78
	}
L75:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v326+(v322^int32(-1))<<(uint(int32(2))%32))))
	v340 = v332
	goto L74
L76:
	;
	goto L77
L77:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v340 = v334 + v322<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L78:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v340)+16)))
	v345 = v340 + v344
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v345)+12)))
	v348 = v346 & int32(_a_F_gist_redo_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v345)+12)) = uint16(v348)
	*(*int64)(unsafe.Add(mBase, uint32(v340))) = base.I64_rotl(v290, int64(32))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	F_MarkBufferDirty(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	goto L73
L80:
	;
	F_UnlockReleaseBuffer(m, v359)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	goto L4
L82:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v368)+16))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+24)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = v371
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v368)))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+64)) = v373
	F_ResolveRecoveryConflictWithSnapshotFullXid(m, v369, v370, v25-int32(-64))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	goto L4
L84:
	;
	v381 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v389 = v2
	v394 = v2
	v400 = v2
	goto L87
L85:
	;
	v738 = v29
	v747 = v2
	goto L86
L86:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v738)+72))
	if v759 < int32(0) {
		goto L145
	} else {
		goto L146
	}
L87:
	;
	v407 = v389 + int32(1)
	v409 = v407 & int32(255)
	v410 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v409, v410, v410, v25+int32(92))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L10
	} else {
		goto L89
	}
L88:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v738 = v736
	v747 = v733
	goto L86
L89:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	v417 = F_XLogInitBufferForRedo(m, l0, v409)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L10
	} else {
		goto L91
	}
L90:
	;
	v437 = int32(0)
	v440 = v25 + int32(76)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+72))
	if v443 < v409 {
		v465 = v437
		goto L96
	} else {
		goto L97
	}
L91:
	;
	if v417 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v422+(v417^int32(-1))<<(uint(int32(2))%32))))
	v436 = v428
	goto L90
L93:
	;
	goto L94
L94:
	;
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v436 = v430 + v417<<(uint(int32(13))%32) + int32(-8192)
	goto L90
L95:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v472 = F_palloc(m, v469<<(uint(int32(2))%32))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L10
	} else {
		goto L106
	}
L96:
	;
	v468 = v465
	goto L95
L97:
	;
	v447 = v442 + v409*int32(52)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+76)))
	if v448 != int32(1) {
		v465 = v437
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v452 = v447 + int32(76)
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452)+43)))
	if v453 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v440 == int32(0) {
		v465 = v437
		goto L96
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v440 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v458 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v458
	v468 = v458
	goto L95
L103:
	;
	v461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = v461
	goto L105
L104:
	;
	goto L105
L105:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v452)+44))
	v465 = v463
	goto L96
L106:
	;
	if v469 <= int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v619 = v400 | base.B2i32(v416 == v437)
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+16)))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	v622 = int32(0)
	v624 = v620 & base.B2i32(v621 != v622)
	if v417 < v622 {
		goto L122
	} else {
		goto L123
	}
L108:
	;
	v477 = v469 & int32(3)
	v478 = int32(4)
	v479 = v468 + v478
	if base.Ui32(v469) < base.Ui32(v478) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v563 = v541
	v568 = v546
	v576 = int32(0)
	goto L117
L110:
	;
	v541 = v479
	v546 = int32(0)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v486 = int32(0)
	v489 = v479
	v494 = v486
	v503 = v486
	goto L113
L113:
	;
	v512 = v472 + v494<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v512))) = v489
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v489)+6)))
	v515 = int32(_a_F_gist_redo_1)
	v517 = v489 + v514&v515
	*(*int32)(unsafe.Add(mBase, uint32(v512)+4)) = v517
	v519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v517)+6)))
	v522 = v517 + v519&v515
	*(*int32)(unsafe.Add(mBase, uint32(v512)+8)) = v522
	v524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v522)+6)))
	v527 = v522 + v524&v515
	*(*int32)(unsafe.Add(mBase, uint32(v512)+12)) = v527
	v529 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v527)+6)))
	v532 = v527 + v529&v515
	v533 = int32(4)
	v534 = v494 + v533
	v536 = v503 + v533
	if v536 != v469&int32(2147483644) {
		v489 = v532
		v494 = v534
		v503 = v536
		goto L113
	} else {
		goto L115
	}
L114:
	;
	if v477 == int32(0) {
		goto L107
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	v541 = v532
	v546 = v534
	goto L109
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v472+v568<<(uint(int32(2))%32)))) = v563
	v588 = int32(1)
	v590 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563)+6)))
	v595 = v576 + v588
	if v595 != v477 {
		v563 = v563 + v590&int32(_a_F_gist_redo_1)
		v568 = v568 + v588
		v576 = v595
		goto L117
	} else {
		goto L119
	}
L118:
	;
	goto L107
L119:
	;
	goto L118
L120:
	;
	F_gistfillbuffer(m, v436, v472, v469, int32(1))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L10
	} else {
		goto L125
	}
L121:
	;
	F_PageInit(m, v642, int32(_a_F_gist_redo_6), int32(16))
	mBase = m.M
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v642)+16)))
	v647 = v642 + v646
	v648 = int32(_a_F_gist_redo_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v647)+14)) = uint16(v648)
	*(*uint16)(unsafe.Add(mBase, uint32(v647)+12)) = uint16(v624)
	*(*int32)(unsafe.Add(mBase, uint32(v647)+8)) = int32(-1)
	goto L120
L122:
	;
	v628 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v628+(v417^int32(-1))<<(uint(int32(2))%32))))
	v642 = v634
	goto L121
L123:
	;
	goto L124
L124:
	;
	v636 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v642 = v636 + v417<<(uint(int32(13))%32) + int32(-8192)
	goto L121
L125:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	if v656 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v436))) = base.I64_rotl(v381, int64(32))
	F_MarkBufferDirty(m, v417)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L10
	} else {
		goto L138
	}
L127:
	;
	v659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v436)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v436+v659)+8)) = int32(-1)
	v663 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v436)+16)))
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v379)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v436+v663))) = base.I64_rotl(v665, int64(32))
	v669 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v436)+16)))
	v670 = v436 + v669
	v671 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670)+12)))
	v673 = v671 & int32(_a_F_gist_redo_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v670)+12)) = uint16(v673)
	goto L126
L128:
	;
	goto L129
L129:
	;
	v675 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379)+18)))
	if v389 < v675-int32(1) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v697 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v436)+16)))
	v699 = *(*int64)(unsafe.Add(mBase, uint32(v379)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v436+v697))) = base.I64_rotl(v699, int64(32))
	v703 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379)+18)))
	v704 = int32(1)
	if (base.B2i32(v703-v704 <= v389)|v619)&v704 != 0 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	v683 = int32(0)
	F_XLogRecGetBlockTag(m, l0, (v389+int32(2))&int32(255), v683, v683, v25+int32(88))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L10
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v693 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v436)+16)))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	*(*int32)(unsafe.Add(mBase, uint32(v436+v693)+8)) = v695
	goto L130
L134:
	;
	v689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v436)+16)))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v25)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v436+v689)+8)) = v691
	goto L130
L135:
	;
	v719 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v436)+16)))
	v720 = v436 + v719
	v721 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v720)+12)))
	v723 = v721 & int32(_a_F_gist_redo_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v720)+12)) = uint16(v723)
	goto L126
L136:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+20)))
	if v710 != int32(1) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v713 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v436)+16)))
	v714 = v436 + v713
	v715 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v714)+12)))
	v717 = v715 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v714)+12)) = uint16(v717)
	goto L126
L138:
	;
	if v389 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379)+18)))
	if base.Ui32(v407) < base.Ui32(v734) {
		v389 = v407
		v394 = v733
		v400 = v619
		goto L87
	} else {
		goto L144
	}
L140:
	;
	v733 = v417
	goto L139
L141:
	;
	goto L142
L142:
	;
	F_UnlockReleaseBuffer(m, v417)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L10
	} else {
		goto L143
	}
L143:
	;
	v733 = v394
	goto L139
L144:
	;
	goto L88
L145:
	;
	F_UnlockReleaseBuffer(m, v747)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L10
	} else {
		goto L159
	}
L146:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738)+76)))
	if v762 != int32(1) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v765 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v769 = F_XLogReadBufferForRedo(m, l0, int32(0), v25+int32(76))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	if v769&int32(-3) == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if v775 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	goto L151
L151:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if v812 == int32(0) {
		goto L145
	} else {
		goto L157
	}
L152:
	;
	v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v793)+16)))
	v797 = base.I64_rotl(v765, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v794+v793))) = v797
	v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v793)+16)))
	v800 = v793 + v799
	v801 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v800)+12)))
	v803 = v801 & int32(_a_F_gist_redo_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v800)+12)) = uint16(v803)
	*(*int64)(unsafe.Add(mBase, uint32(v793))) = v797
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	F_MarkBufferDirty(m, v806)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L10
	} else {
		goto L156
	}
L153:
	;
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v779+(v775^int32(-1))<<(uint(int32(2))%32))))
	v793 = v785
	goto L152
L154:
	;
	goto L155
L155:
	;
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v793 = v787 + v775<<(uint(int32(13))%32) + int32(-8192)
	goto L152
L156:
	;
	goto L151
L157:
	;
	F_UnlockReleaseBuffer(m, v812)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L10
	} else {
		goto L158
	}
L158:
	;
	goto L145
L159:
	;
	goto L4
L160:
	;
	if v827 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v831 = *(*int64)(unsafe.Add(mBase, uint32(v822)))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	if v832 < int32(0) {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	goto L163
L163:
	;
	v872 = F_XLogReadBufferForRedo(m, l0, int32(1), v25+int32(76))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L10
	} else {
		goto L169
	}
L164:
	;
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v850)+16)))
	v852 = v851 + v850
	v853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v852)+12)))
	v855 = v853 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v852)+12)) = uint16(v855)
	*(*int64)(unsafe.Add(mBase, uint32(v850)+24)) = v831
	v858 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v850)+12)) = uint16(v858)
	*(*int64)(unsafe.Add(mBase, uint32(v850))) = base.I64_rotl(v823, int64(32))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	F_MarkBufferDirty(m, v863)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L10
	} else {
		goto L168
	}
L165:
	;
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v836+(v832^int32(-1))<<(uint(int32(2))%32))))
	v850 = v842
	goto L164
L166:
	;
	goto L167
L167:
	;
	v844 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v850 = v844 + v832<<(uint(int32(13))%32) + int32(-8192)
	goto L164
L168:
	;
	goto L163
L169:
	;
	if v872 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if v876 < int32(0) {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	goto L172
L172:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	if v905 != 0 {
		goto L179
	} else {
		goto L180
	}
L173:
	;
	v895 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v822)+8)))
	F_PageIndexTupleDelete(m, v894, v895)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L10
	} else {
		goto L177
	}
L174:
	;
	v880 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[3]))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v880+(v876^int32(-1))<<(uint(int32(2))%32))))
	v894 = v886
	goto L173
L175:
	;
	goto L176
L176:
	;
	v888 = *(*int32)(unsafe.Add(mBase, _c_F_gist_redo[4]))
	v894 = v888 + v876<<(uint(int32(13))%32) + int32(-8192)
	goto L173
L177:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v894))) = base.I64_rotl(v823, int64(32))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	F_MarkBufferDirty(m, v901)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L10
	} else {
		goto L178
	}
L178:
	;
	goto L172
L179:
	;
	F_UnlockReleaseBuffer(m, v905)
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L10
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v25)+92))
	if v908 == int32(0) {
		goto L4
	} else {
		goto L183
	}
L182:
	;
	goto L181
L183:
	;
	F_UnlockReleaseBuffer(m, v908)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L10
	} else {
		goto L184
	}
L184:
	;
	goto L4
L185:
	;
	m.G0 = v25 + int32(96)
	return
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v107
	F_errmsg_internal(m, int32(_a_F_gist_redo_8), v25+int32(32))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L10
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_gist_redo_9), int32(102), int32(_a_F_gist_redo_10))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L10
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v158
	F_errmsg_internal(m, int32(_a_F_gist_redo_8), v25+int32(16))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L10
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_gist_redo_9), int32(135), int32(_a_F_gist_redo_10))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L10
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v36
	F_errmsg_internal(m, int32(_a_F_gist_redo_11), v25)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L10
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_gist_redo_9), int32(422), int32(_a_F_gist_redo_12))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L10
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
