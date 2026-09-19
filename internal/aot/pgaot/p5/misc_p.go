package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ParamsErrorCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	if l0 == int32(0) {
		m.G0 = v6 + int32(32)
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 == int32(0) {
			m.G0 = v6 + int32(32)
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
			if v13 == int32(0) {
				m.G0 = v6 + int32(32)
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v16 == int32(0) {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
						F_errcontext_msg(m, int32(_a_F_ParamsErrorCallback_0), v6)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							m.G0 = v6 + int32(32)
							return
						}
					}
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if v19 == int32(0) {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
							F_errcontext_msg(m, int32(_a_F_ParamsErrorCallback_0), v6)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								m.G0 = v6 + int32(32)
								return
							}
						}
					} else {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v27
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v25
							F_errcontext_msg(m, int32(_a_F_ParamsErrorCallback_1), v6+int32(16))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								m.G0 = v6 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_PartConstraintImpliedByRelConstraint(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v11 == v3 {
		v83 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v87 = F_ConstraintImpliedByRelConstraint(m, l0, l1, v83)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L10
	} else {
		goto L15
	}
L2:
	;
	v14 = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v15 != v14 {
		v83 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v18 <= int32(0) {
		v83 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = v14
	v26 = v3
	goto L5
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = v25 - int32(1)
	v35 = v30 + v32<<(uint(int32(4))%32)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+31)))
	if v36 != int32(118) {
		v73 = v26
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v83 = v73
	goto L1
L7:
	;
	v76 = v25 + int32(1)
	if v76 <= v18 {
		v25 = v76
		v26 = v73
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+29)))
	if v39 != 0 {
		v73 = v26
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v42 = F_palloc0(m, int32(20))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(52)
	v55 = v30 + v40<<(uint(int32(4))%32) + v32*int32(100)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+88))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+96))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+116))
	v60 = F_makeVar(m, int32(1), base.I32_extend16_s(v25), v56, v57, v58, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+16)) = int32(-1)
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+12)) = uint8(v64)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v60
	v69 = F_lappend(m, v26, v42)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v73 = v69
	goto L7
L14:
	;
	goto L6
L15:
	;
	return v87
}
func F_PhysicalReplicationSlotNewXmin(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_PhysicalReplicationSlotNewXmin[0]))
	v10 = base.AtomicRmwXchg32(m, v7, v3, int32(1))
	if v10 != 0 {
		F_s_lock(m, v7, int32(_a_F_PhysicalReplicationSlotNewXmin_0), int32(2527), int32(_a_F_PhysicalReplicationSlotNewXmin_1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_PhysicalReplicationSlotNewXmin[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = int32(0)
			if base.Ui32(l0) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
				v44 = int32(1)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+96))
				if base.Ui32(v22) < base.Ui32(int32(3)) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
					v44 = int32(1)
				} else {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v22)) == int32(0) {
						v36 = base.B2i32(base.Ui32(v22) < base.Ui32(l0))
					} else {
						v36 = int32(base.Ui32(v22-l0) >> (uint(int32(31)) % 32))
					}
					if v36 == int32(0) {
						v44 = v3
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
						v44 = int32(1)
					}
				}
			}
			if base.Ui32(l1) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
				v67 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v67))
				F_ReplicationSlotMarkDirty(m)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+100))
				if base.Ui32(v47) < base.Ui32(int32(3)) {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
					v67 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v67))
					F_ReplicationSlotMarkDirty(m)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v47)) == int32(0) {
						v61 = base.B2i32(base.Ui32(v47) < base.Ui32(l1))
					} else {
						v61 = int32(base.Ui32(v47-l1) >> (uint(int32(31)) % 32))
					}
					if v61 == int32(0) {
						v70 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v70))
						if v44 == v70 {
							return
						} else {
							F_ReplicationSlotMarkDirty(m)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
						v67 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v67))
						F_ReplicationSlotMarkDirty(m)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_PhysicalReplicationSlotNewXmin[1]))
		*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = int32(0)
		if base.Ui32(l0) < base.Ui32(int32(3)) {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
			v44 = int32(1)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+96))
			if base.Ui32(v22) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
				v44 = int32(1)
			} else {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v22)) == int32(0) {
					v36 = base.B2i32(base.Ui32(v22) < base.Ui32(l0))
				} else {
					v36 = int32(base.Ui32(v22-l0) >> (uint(int32(31)) % 32))
				}
				if v36 == int32(0) {
					v44 = v3
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = l0
					v44 = int32(1)
				}
			}
		}
		if base.Ui32(l1) < base.Ui32(int32(3)) {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
			v67 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v67))
			F_ReplicationSlotMarkDirty(m)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+100))
			if base.Ui32(v47) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
				v67 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v67))
				F_ReplicationSlotMarkDirty(m)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v47)) == int32(0) {
					v61 = base.B2i32(base.Ui32(v47) < base.Ui32(l1))
				} else {
					v61 = int32(base.Ui32(v47-l1) >> (uint(int32(31)) % 32))
				}
				if v61 == int32(0) {
					v70 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v70))
					if v44 == v70 {
						return
					} else {
						F_ReplicationSlotMarkDirty(m)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v7)+100)) = l1
					v67 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v7))), uint32(v67))
					F_ReplicationSlotMarkDirty(m)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							return
						}
					}
				}
			}
		}
	}
}
func F_PopActiveSnapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v8 - int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v13 != 0 {
		v19 = v5
		F_pfree(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0])) = v6
			if v6 != 0 {
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
				if v26 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
					v33 = v31 - int32(48)
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v34))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v29)) == int32(0) {
						v46 = base.B2i32(base.Ui32(v29) < base.Ui32(v34))
					} else {
						v46 = int32(base.Ui32(v29-v34) >> (uint(int32(31)) % 32))
					}
					if v46 == int32(0) {
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v50 = v49
						*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
						v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
					}
				} else {
					v50 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
					v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
				}
			}
			return
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
		if v14 != 0 {
			v19 = v5
			F_pfree(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0])) = v6
				if v6 != 0 {
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
					if v26 != 0 {
						v28 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
						v33 = v31 - int32(48)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v34))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v29)) == int32(0) {
							v46 = base.B2i32(base.Ui32(v29) < base.Ui32(v34))
						} else {
							v46 = int32(base.Ui32(v29-v34) >> (uint(int32(31)) % 32))
						}
						if v46 == int32(0) {
						} else {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							v50 = v49
							*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
						}
					} else {
						v50 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
						v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
					}
				}
				return
			}
		} else {
			F_pfree(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0]))
				v19 = v18
				F_pfree(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0])) = v6
					if v6 != 0 {
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
						if v26 != 0 {
							v28 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
							v31 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
							v33 = v31 - int32(48)
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v34))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v29)) == int32(0) {
								v46 = base.B2i32(base.Ui32(v29) < base.Ui32(v34))
							} else {
								v46 = int32(base.Ui32(v29-v34) >> (uint(int32(31)) % 32))
							}
							if v46 == int32(0) {
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
								v50 = v49
								*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
								v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
							}
						} else {
							v50 = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3])) = v50
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v50
						}
					}
					return
				}
			}
		}
	}
}
func F_PostgresSendReadyForQueryIfNecessary(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])))
	if v4 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[1]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
		if (v9-int32(7))&int32(-9) == int32(0) {
			v17 = int32(0)
			F_pgstat_report_activity(m, int32(6), v17)
			mBase = m.M
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[2]))
			if v20 <= v17 {
				F_ReportChangedGUCOptions(m)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return
				} else {
					v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
					F_ReadyForQuery(m, v97)
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return
					} else {
						v101 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
						return
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[4]))
				if v24 != 0 {
					v27 = base.B2i32(v24 <= v20)
				} else {
					v27 = int32(0)
				}
				if v27 != 0 {
					F_ReportChangedGUCOptions(m)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
						F_ReadyForQuery(m, v97)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v101 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
							return
						}
					}
				} else {
					v29 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[5])) = uint8(v29)
					F_enable_timeout_after(m, int32(7), v20)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_ReportChangedGUCOptions(m)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
							F_ReadyForQuery(m, v97)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v101 = int32(0)
								*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
								return
							}
						}
					}
				}
			}
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[1]))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
			if v36 != int32(0) {
				v40 = int32(0)
				F_pgstat_report_activity(m, int32(4), v40)
				mBase = m.M
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[2]))
				if v43 <= v40 {
					F_ReportChangedGUCOptions(m)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
						F_ReadyForQuery(m, v97)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v101 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
							return
						}
					}
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[4]))
					if v47 != 0 {
						v50 = base.B2i32(v47 <= v43)
					} else {
						v50 = int32(0)
					}
					if v50 != 0 {
						F_ReportChangedGUCOptions(m)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
							F_ReadyForQuery(m, v97)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v101 = int32(0)
								*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
								return
							}
						}
					} else {
						v52 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[5])) = uint8(v52)
						F_enable_timeout_after(m, int32(7), v43)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_ReportChangedGUCOptions(m)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
								F_ReadyForQuery(m, v97)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v101 = int32(0)
									*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
									return
								}
							}
						}
					}
				}
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[6]))
				if v58 != 0 {
					F_ProcessNotifyInterrupt(m, int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v63 = F_pgstat_report_stat(m, int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[7])))
							if int32(0) < v63 {
								if v68 != 0 {
									v80 = int32(0)
									F_pgstat_report_activity(m, int32(2), v80)
									mBase = m.M
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
									if v83 <= v80 {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v97)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v101 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
												return
											}
										}
									} else {
										v87 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v87)
										F_enable_timeout_after(m, int32(9), v83)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													v101 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
													return
												}
											}
										}
									}
								} else {
									F_enable_timeout_after(m, int32(10), v63)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										v80 = int32(0)
										F_pgstat_report_activity(m, int32(2), v80)
										mBase = m.M
										v83 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
										if v83 <= v80 {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													v101 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
													return
												}
											}
										} else {
											v87 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v87)
											F_enable_timeout_after(m, int32(9), v83)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												F_ReportChangedGUCOptions(m)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
													F_ReadyForQuery(m, v97)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														v101 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
														return
													}
												}
											}
										}
									}
								}
							} else {
								if v68 == int32(0) {
									v80 = int32(0)
									F_pgstat_report_activity(m, int32(2), v80)
									mBase = m.M
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
									if v83 <= v80 {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v97)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v101 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
												return
											}
										}
									} else {
										v87 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v87)
										F_enable_timeout_after(m, int32(9), v83)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													v101 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
													return
												}
											}
										}
									}
								} else {
									F_disable_timeout(m, int32(10))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v80 = int32(0)
										F_pgstat_report_activity(m, int32(2), v80)
										mBase = m.M
										v83 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
										if v83 <= v80 {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													v101 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
													return
												}
											}
										} else {
											v87 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v87)
											F_enable_timeout_after(m, int32(9), v83)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												F_ReportChangedGUCOptions(m)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
													F_ReadyForQuery(m, v97)
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														v101 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
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
				} else {
					v63 = F_pgstat_report_stat(m, int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[7])))
						if int32(0) < v63 {
							if v68 != 0 {
								v80 = int32(0)
								F_pgstat_report_activity(m, int32(2), v80)
								mBase = m.M
								v83 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
								if v83 <= v80 {
									F_ReportChangedGUCOptions(m)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
										F_ReadyForQuery(m, v97)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v101 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
											return
										}
									}
								} else {
									v87 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v87)
									F_enable_timeout_after(m, int32(9), v83)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v97)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v101 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
												return
											}
										}
									}
								}
							} else {
								F_enable_timeout_after(m, int32(10), v63)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									v80 = int32(0)
									F_pgstat_report_activity(m, int32(2), v80)
									mBase = m.M
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
									if v83 <= v80 {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v97)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v101 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
												return
											}
										}
									} else {
										v87 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v87)
										F_enable_timeout_after(m, int32(9), v83)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													v101 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
													return
												}
											}
										}
									}
								}
							}
						} else {
							if v68 == int32(0) {
								v80 = int32(0)
								F_pgstat_report_activity(m, int32(2), v80)
								mBase = m.M
								v83 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
								if v83 <= v80 {
									F_ReportChangedGUCOptions(m)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
										F_ReadyForQuery(m, v97)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v101 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
											return
										}
									}
								} else {
									v87 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v87)
									F_enable_timeout_after(m, int32(9), v83)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v97)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v101 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
												return
											}
										}
									}
								}
							} else {
								F_disable_timeout(m, int32(10))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return
								} else {
									v80 = int32(0)
									F_pgstat_report_activity(m, int32(2), v80)
									mBase = m.M
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[8]))
									if v83 <= v80 {
										F_ReportChangedGUCOptions(m)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
											F_ReadyForQuery(m, v97)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												v101 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
												return
											}
										}
									} else {
										v87 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[9])) = uint8(v87)
										F_enable_timeout_after(m, int32(9), v83)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											F_ReportChangedGUCOptions(m)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[3]))
												F_ReadyForQuery(m, v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													v101 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_PostgresSendReadyForQueryIfNecessary[0])) = uint8(v101)
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
	} else {
		return
	}
}
func F_PreventCommandDuringRecovery(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PreventCommandDuringRecovery[0])))
	if v9 == int32(1) {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_PreventCommandDuringRecovery[1]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+316))
		v17 = base.B2i32(v15 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_PreventCommandDuringRecovery[0])) = uint8(v17)
		v19 = v17
	} else {
		v19 = int32(0)
	}
	if v19 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(100663618))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(_a_F_PreventCommandDuringRecovery_0), v5)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PreventCommandDuringRecovery_1), int32(448), int32(_a_F_PreventCommandDuringRecovery_2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_ProcessCommittedInvalidationMessages(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l1 <= int32(0) {
		m.G0 = v9 + int32(32)
		return
	} else {
		v15 = F_errstart(m, int32(11), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
				if l2 != 0 {
					v20 = int32(_a_F_ProcessCommittedInvalidationMessages_0)
				} else {
					v20 = int32(_a_F_ProcessCommittedInvalidationMessages_1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v20
				F_errmsg_internal(m, int32(_a_F_ProcessCommittedInvalidationMessages_2), v9+int32(16))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ProcessCommittedInvalidationMessages_3), int32(1143), int32(_a_F_ProcessCommittedInvalidationMessages_4))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						if l2 != 0 {
							v34 = F_errstart(m, int32(11), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								if v34 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
									F_errmsg_internal(m, int32(_a_F_ProcessCommittedInvalidationMessages_5), v9)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ProcessCommittedInvalidationMessages_3), int32(1147), int32(_a_F_ProcessCommittedInvalidationMessages_4))
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return
										} else {
											if l3 != 0 {
												v46 = F_GetDatabasePath(m, l3, l4)
												mBase = m.M
												v47 = m.ExcPending
												if v47 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
													F_RelationCacheInitFilePreInvalidate(m)
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return
													} else {
														v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
														F_pfree(m, v52)
														mBase = m.M
														v54 = m.ExcPending
														if v54 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
															F_SIInsertDataEntries(m, l0, l1)
															mBase = m.M
															v61 = m.ExcPending
															if v61 != 0 {
																return
															} else {
																F_RelationCacheInitFilePostInvalidate(m)
																mBase = m.M
																v63 = m.ExcPending
																if v63 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													}
												}
											} else {
												F_RelationCacheInitFilePreInvalidate(m)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													F_SIInsertDataEntries(m, l0, l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_RelationCacheInitFilePostInvalidate(m)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
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
								} else {
									if l3 != 0 {
										v46 = F_GetDatabasePath(m, l3, l4)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
											F_RelationCacheInitFilePreInvalidate(m)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
												F_pfree(m, v52)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
													F_SIInsertDataEntries(m, l0, l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_RelationCacheInitFilePostInvalidate(m)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									} else {
										F_RelationCacheInitFilePreInvalidate(m)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_SIInsertDataEntries(m, l0, l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_RelationCacheInitFilePostInvalidate(m)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
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
						} else {
							F_SIInsertDataEntries(m, l0, l1)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			} else {
				if l2 != 0 {
					v34 = F_errstart(m, int32(11), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						if v34 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
							F_errmsg_internal(m, int32(_a_F_ProcessCommittedInvalidationMessages_5), v9)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ProcessCommittedInvalidationMessages_3), int32(1147), int32(_a_F_ProcessCommittedInvalidationMessages_4))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									if l3 != 0 {
										v46 = F_GetDatabasePath(m, l3, l4)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
											F_RelationCacheInitFilePreInvalidate(m)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
												F_pfree(m, v52)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
													F_SIInsertDataEntries(m, l0, l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_RelationCacheInitFilePostInvalidate(m)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									} else {
										F_RelationCacheInitFilePreInvalidate(m)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_SIInsertDataEntries(m, l0, l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_RelationCacheInitFilePostInvalidate(m)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
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
						} else {
							if l3 != 0 {
								v46 = F_GetDatabasePath(m, l3, l4)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
									F_RelationCacheInitFilePreInvalidate(m)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
										F_pfree(m, v52)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
											F_SIInsertDataEntries(m, l0, l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_RelationCacheInitFilePostInvalidate(m)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							} else {
								F_RelationCacheInitFilePreInvalidate(m)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_SIInsertDataEntries(m, l0, l1)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										F_RelationCacheInitFilePostInvalidate(m)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
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
				} else {
					F_SIInsertDataEntries(m, l0, l1)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
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
func F_ProcessInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int64
	_ = v325
	var v329 int64
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v617 int32
	_ = v617
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int64
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v790 int32
	_ = v790
	var v791 int64
	_ = v791
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	if v17 != 0 {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L18
	} else {
		goto L361
	}
L2:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L18
	} else {
		goto L356
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L18
	} else {
		goto L352
	}
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L18
	} else {
		goto L348
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1])) = int32(0)
	F_LockErrorCleanup(m)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L18
	} else {
		goto L343
	}
L6:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L18
	} else {
		goto L338
	}
L7:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L18
	} else {
		goto L333
	}
L8:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L18
	} else {
		goto L328
	}
L9:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L18
	} else {
		goto L324
	}
L10:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L18
	} else {
		goto L320
	}
L11:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L18
	} else {
		goto L316
	}
L12:
	;
	m.G0 = v14 + int32(32)
	return
L13:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2]))
	if v19 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[3])) = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[4]))
	if v24 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v26 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[4])) = v26
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1])) = v26
	F_LockErrorCleanup(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[5]))
	if v149 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L18:
	;
	return
L19:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[6])))
	if v34 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v34 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[7]))
	if v38 != int32(2) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[7])) = int32(0)
	goto L1
L23:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[8]))
	if v45 == int32(4) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[9]))
	if v49 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[10]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[11]))
	if v52 == v54 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v58 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L18
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[8]))
	if v73 != int32(12) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	if v58 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_0), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L18
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L18
	} else {
		goto L35
	}
L33:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3435), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L18
	} else {
		goto L57
	}
L37:
	;
	switch v73 - int32(5) {
	case 0:
		goto L40
	default:
		goto L36
	case 9:
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v118 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L18
	} else {
		goto L50
	}
L40:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L18
	} else {
		goto L46
	}
L41:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_3), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3446), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L18
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
	F_errcode(m, int32(16908741))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v102 + int32(96)
	F_errmsg(m, int32(_a_F_ProcessInterrupts_4), v14+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3451), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	if v118 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_5), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L18
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L18
	} else {
		goto L56
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3455), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L18
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_6), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L18
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3462), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L18
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[13]))
	if v270 != 0 {
		goto L5
	} else {
		goto L85
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[5])) = int32(0)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[14])))
	if v156 != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[15]))
	if v158 <= int32(0) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v161 = m.G0
	v163 = v161 - int32(48)
	m.G0 = v163
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[16]))
	v167 = int32(0)
	F_ModifyWaitEvent(m, v166, v167, int32(128), v167)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[16]))
	v175 = int32(0)
	v178 = F_WaitEventSetWait(m, v174, v175, v163, int32(3), v175)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L18
	} else {
		goto L67
	}
L66:
	;
	m.G0 = v163 + int32(48)
	if v234 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L67:
	;
	if v178 <= int32(0) {
		v234 = int32(1)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v185 = v178
	goto L69
L69:
	;
	v195 = int32(0)
	goto L71
L70:
	;
	v234 = int32(1)
	goto L66
L71:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v163+v195<<(uint(int32(4))%32))+4))
	v210 = v208 & int32(128)
	v212 = base.B2i32(v210 == int32(0))
	if v210 != 0 {
		v234 = v212
		goto L66
	} else {
		goto L73
	}
L72:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = int32(0)
	goto L78
L73:
	;
	if v208&int32(1) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v218 = v195 + int32(1)
	if v218 == v185 {
		v234 = v212
		goto L66
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	goto L72
L77:
	;
	v195 = v218
	goto L71
L78:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[16]))
	v227 = int32(0)
	v230 = F_WaitEventSetWait(m, v226, v227, v163, int32(3), v227)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L18
	} else {
		goto L79
	}
L79:
	;
	if int32(0) < v230 {
		v185 = v230
		goto L69
	} else {
		goto L80
	}
L80:
	;
	goto L70
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[13])) = int32(1)
	goto L61
L82:
	;
	goto L83
L83:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[15]))
	F_enable_timeout_after(m, int32(11), v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	goto L61
L85:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1]))
	if v272 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[18]))
	if v342 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L87:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1]))
	if v283 == int32(0) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[19]))
	if v276 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[3])) = int32(1)
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1])) = int32(0)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[20])))
	if v293&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	if v316 != 0 {
		goto L6
	} else {
		goto L107
	}
L92:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[21])))
	if v308&int32(1) != 0 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v298 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[20])) = uint8(v298)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v301 = v293 & int32(1)
	goto L92
L96:
	;
	v317 = int32(0)
	if base.B2i32(v301 == int32(0))|base.B2i32(v316 == v317) == v317 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v313 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[21])) = uint8(v313)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v316 = v308 & int32(1)
	goto L96
L100:
	;
	v325 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessInterrupts[22]))
	goto L103
L101:
	;
	goto L102
L102:
	;
	if v301 != 0 {
		goto L2
	} else {
		goto L106
	}
L103:
	;
	v329 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessInterrupts[23]))
	goto L104
L104:
	;
	if v325 < v329 {
		goto L91
	} else {
		goto L105
	}
L105:
	;
	goto L2
L106:
	;
	goto L91
L107:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[8]))
	if v332 == int32(4) {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[14])))
	if v336 == int32(0) {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	goto L86
L110:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[24]))
	if v516 != 0 {
		goto L163
	} else {
		goto L164
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[18])) = int32(0)
	v349 = int32(7)
	goto L112
L112:
	;
	v361 = v349 << (uint(int32(2)) % 32)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+uint32(_c_F_ProcessInterrupts[25])))
	if v362 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L110
L114:
	;
	v501 = v349 + int32(1)
	if v501 != int32(14) {
		v349 = v501
		goto L112
	} else {
		goto L162
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361)+uint32(_c_F_ProcessInterrupts[25]))) = int32(0)
	switch v349 - int32(7) {
	case 0:
		goto L116
	case 1, 2, 3:
		goto L118
	case 4:
		goto L117
	case 5:
		goto L120
	case 6:
		goto L122
	default:
		goto L121
	}
L116:
	;
	F_pgstat_report_recovery_conflict(m, v349)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L18
	} else {
		goto L152
	}
L117:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26]))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+24))
	goto L139
L118:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26]))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+24))
	goto L134
L119:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[27]))
	v412 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v411)+73)) = uint8(v412)
	goto L118
L120:
	;
	v406 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L18
	} else {
		goto L132
	}
L121:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L18
	} else {
		goto L129
	}
L122:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[28]))
	if v372 == int32(0) {
		goto L110
	} else {
		goto L123
	}
L123:
	;
	v375 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L18
	} else {
		goto L124
	}
L124:
	;
	if v375 != 0 {
		goto L119
	} else {
		goto L125
	}
L125:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[29]))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)+72))
	goto L126
L126:
	;
	if int32(0) <= v379 {
		goto L110
	} else {
		goto L127
	}
L127:
	;
	v383 = int32(_a_F_ProcessInterrupts_7)
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[30]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[31])) = int32(1)
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[17]))
	F_SetLatch(m, v389)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[30])) = v384
	goto L128
L128:
	;
	goto L110
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v349
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_8), v14)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L18
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3358), int32(_a_F_ProcessInterrupts_9))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L18
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	if v406 == int32(0) {
		goto L114
	} else {
		goto L133
	}
L133:
	;
	goto L119
L134:
	;
	if base.B2i32(v416 != int32(0)) == int32(0) {
		goto L114
	} else {
		goto L135
	}
L135:
	;
	if v349 == int32(11) {
		goto L117
	} else {
		goto L136
	}
L136:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26]))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)+28))
	goto L137
L137:
	;
	if int32(1) < v425 {
		goto L116
	} else {
		goto L138
	}
L138:
	;
	goto L117
L139:
	;
	if (v430-int32(7))&int32(-9) == int32(0) {
		goto L114
	} else {
		goto L140
	}
L140:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[14])))
	if v438 != 0 {
		goto L116
	} else {
		goto L141
	}
L141:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[19]))
	if v440 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v441 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v361)+uint32(_c_F_ProcessInterrupts[25]))) = v441
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[18])) = v441
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[3])) = v441
	goto L114
L143:
	;
	goto L144
L144:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	F_pgstat_report_recovery_conflict(m, v349)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L18
	} else {
		goto L147
	}
L147:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L18
	} else {
		goto L148
	}
L148:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_10), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errdetail_recovery_conflict(m, v349)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3331), int32(_a_F_ProcessInterrupts_9))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L18
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	if v349 == int32(7) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v481 = int32(67240389)
	goto L156
L155:
	;
	v481 = int32(16777220)
	goto L156
L156:
	;
	F_errcode(m, v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_11), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L18
	} else {
		goto L158
	}
L158:
	;
	F_errdetail_recovery_conflict(m, v349)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L18
	} else {
		goto L159
	}
L159:
	;
	F_errhint(m, int32(_a_F_ProcessInterrupts_12), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L18
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3354), int32(_a_F_ProcessInterrupts_9))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L18
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	goto L113
L163:
	;
	v518 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[24])) = v518
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[32]))
	if v518 < v521 {
		goto L11
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[33]))
	if v525 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L165
L167:
	;
	v527 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[33])) = v527
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[34]))
	if v527 < v530 {
		goto L10
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[35]))
	if v534 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	goto L169
L171:
	;
	v536 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[35])) = v536
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[36]))
	if v536 < v539 {
		goto L9
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[37]))
	v544 = int32(0)
	v547 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[14])))
	if base.B2i32(v543 == v544)|base.B2i32(v547&int32(1) == v544) != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L173
L175:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[38]))
	if v565 != 0 {
		goto L180
	} else {
		goto L181
	}
L176:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26]))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v554)+24))
	goto L177
L177:
	;
	if v555 != int32(0) {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[37])) = int32(0)
	v562 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L18
	} else {
		goto L179
	}
L179:
	;
	goto L175
L180:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L18
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[39]))
	if v569 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	goto L182
L184:
	;
	v570 = m.G0
	v572 = v570 - int32(160)
	m.G0 = v572
	v574 = int32(_a_F_ProcessInterrupts_13)
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v576 + int32(1)
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40]))
	if v581 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	goto L186
L186:
	;
	v917 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41]))
	if v917 != 0 {
		goto L263
	} else {
		goto L264
	}
L187:
	;
	v599 = int32(_a_F_ProcessInterrupts_14)
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[42]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[42])) = v598
	v604 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[39])) = v604
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[43]))
	if base.B2i32(v607 == v604)|base.B2i32(v607 == int32(_a_F_ProcessInterrupts_15)) == v604 {
		goto L193
	} else {
		goto L194
	}
L188:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[44]))
	v591 = F_AllocSetContextCreateInternal(m, v586, int32(_a_F_ProcessInterrupts_16), int32(0), int32(_a_F_ProcessInterrupts_17), int32(_a_F_ProcessInterrupts_18))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L18
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	F_MemoryContextReset(m, v581)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L18
	} else {
		goto L192
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40])) = v591
	v598 = v591
	goto L187
L192:
	;
	v597 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40]))
	v598 = v597
	goto L187
L193:
	;
	v617 = v607
	goto L196
L194:
	;
	v882 = v598
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[42])) = v600
	F_MemoryContextReset(m, v882)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L18
	} else {
		goto L262
	}
L196:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v617)+56))
	if v626 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v880 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40]))
	v882 = v880
	goto L195
L198:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
	if v876 != int32(_a_F_ProcessInterrupts_15) {
		v617 = v876
		goto L196
	} else {
		goto L261
	}
L199:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v617)+20))
	if v629 <= int32(0) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v636 = int32(0)
	goto L201
L201:
	;
	v645 = v636 << (uint(int32(3)) % 32)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v617)+56))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v645+v646)+4))
	if v648 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	goto L198
L203:
	;
	v862 = v636 + int32(1)
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v617)+20))
	if v862 < v863 {
		v636 = v862
		goto L201
	} else {
		goto L260
	}
L204:
	;
	v652 = v648
	goto L205
L205:
	;
	v667 = F_shm_mq_receive(m, v652, v572+int32(56), v572+int32(52), int32(1))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L18
	} else {
		goto L207
	}
L206:
	;
	goto L203
L207:
	;
	if v667 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	if v667 == int32(1) {
		goto L203
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v688 = v572 + int32(36)
	F_initStringInfo(m, v688)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L18
	} else {
		goto L216
	}
L211:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L18
	} else {
		goto L212
	}
L212:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L18
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_19), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L18
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_20), int32(1127), int32(_a_F_ProcessInterrupts_16))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L18
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v572)+52))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v572)+56))
	F_appendBinaryStringInfo(m, v688, v691, v692)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L18
	} else {
		goto L217
	}
L217:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v617)+64))
	if v695 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v709 = F_pq_getmsgbyte(m, v572+int32(36))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L18
	} else {
		goto L227
	}
L219:
	;
	v698 = v695 + v636
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698))))
	if v699 != 0 {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v700 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v698))) = uint8(v700)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v617)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v617)+60)) = v702 + v700
	goto L218
L221:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v572)+36))
	F_pfree(m, v844)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L18
	} else {
		goto L258
	}
L222:
	;
	v830 = v572 + int32(36)
	v832 = F_pq_getmsgint(m, v830, int32(4))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L18
	} else {
		goto L253
	}
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L18
	} else {
		goto L250
	}
L224:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v617)+56))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v805+v645)+4))
	F_shm_mq_detach(m, v807)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L18
	} else {
		goto L249
	}
L225:
	;
	v757 = v572 + int32(36)
	v759 = F_pq_getmsgint(m, v757, int32(4))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L18
	} else {
		goto L242
	}
L226:
	;
	F_pq_parse_errornotice(m, v572+int32(36), v572+int32(60))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L18
	} else {
		goto L228
	}
L227:
	;
	v711 = base.I32_extend8_s(v709)
	switch v711 - int32(65) {
	case 0:
		goto L222
	default:
		goto L223
	case 4, 13:
		goto L226
	case 15:
		goto L225
	case 23:
		goto L224
	}
L228:
	;
	v720 = int32(21)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v572)+60))
	if v720 <= v721 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v724 = v720
	goto L231
L230:
	;
	v724 = v721
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+60)) = v724
	v727 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[45]))
	if v727 != int32(2) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v572)+108))
	if v730 != 0 {
		goto L236
	} else {
		goto L237
	}
L233:
	;
	goto L234
L234:
	;
	v745 = int32(_a_F_ProcessInterrupts_21)
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[46]))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v617)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[46])) = v748
	F_ThrowErrorData(m, v572+int32(60))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L18
	} else {
		goto L241
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+108)) = v742
	goto L234
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+20)) = int32(_a_F_ProcessInterrupts_22)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+16)) = v730
	v737 = F_psprintf(m, int32(_a_F_ProcessInterrupts_23), v572+int32(16))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L18
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v740 = F_pstrdup(m, int32(_a_F_ProcessInterrupts_22))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L18
	} else {
		goto L240
	}
L239:
	;
	v742 = v737
	goto L235
L240:
	;
	v742 = v740
	goto L235
L241:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[46])) = v746
	goto L221
L242:
	;
	v761 = F_pq_getmsgint64(m, v757)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L18
	} else {
		goto L243
	}
L243:
	;
	F_pq_getmsgend(m, v757)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L18
	} else {
		goto L244
	}
L244:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[47]))
	if v767 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	goto L221
L246:
	;
	goto L245
L247:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[48])))
	if v771&int32(1) == int32(0) {
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v776 = int32(_a_F_ProcessInterrupts_24)
	v778 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2]))
	v779 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2])) = v778 + v779
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = v782 + v779
	v790 = v767 + v759<<(uint(int32(3))%32) + int32(232)
	v791 = *(*int64)(unsafe.Add(mBase, uint32(v790)))
	*(*int64)(unsafe.Add(mBase, uint32(v790))) = v791 + v761
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = v794 + v779
	v800 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2])) = v800 - v779
	goto L246
L249:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v617)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v810+v645)+4)) = int32(0)
	goto L221
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = v711
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v572)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+4)) = v819
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_25), v572)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L18
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_20), int32(1249), int32(_a_F_ProcessInterrupts_26))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L18
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	v834 = F_pq_getmsgrawstring(m, v830)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L18
	} else {
		goto L254
	}
L254:
	;
	v836 = F_pq_getmsgrawstring(m, v830)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L18
	} else {
		goto L255
	}
L255:
	;
	F_pq_endmessage(m, v830)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L18
	} else {
		goto L256
	}
L256:
	;
	F_NotifyMyFrontEnd(m, v834, v836, v832)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L18
	} else {
		goto L257
	}
L257:
	;
	goto L221
L258:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v617)+56))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v847+v645)+4))
	if v849 != 0 {
		v652 = v849
		goto L205
	} else {
		goto L259
	}
L259:
	;
	goto L206
L260:
	;
	goto L202
L261:
	;
	goto L197
L262:
	;
	v896 = int32(_a_F_ProcessInterrupts_13)
	v898 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v898 - int32(1)
	m.G0 = v572 + int32(160)
	goto L186
L263:
	;
	F_ProcessLogMemoryContextInterrupt(m)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L18
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v921 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[49]))
	if v921 == int32(0) {
		goto L12
	} else {
		goto L267
	}
L266:
	;
	goto L265
L267:
	;
	v924 = m.G0
	v926 = v924 - int32(176)
	m.G0 = v926
	v928 = int32(_a_F_ProcessInterrupts_13)
	v930 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v930 + int32(1)
	v935 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[50]))
	if v935 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	v953 = int32(0)
	v954 = int32(_a_F_ProcessInterrupts_14)
	v955 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[42]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[42])) = v952
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[49])) = v953
	v962 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[51]))
	if v962 != 0 {
		goto L274
	} else {
		goto L275
	}
L269:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[44]))
	v945 = F_AllocSetContextCreateInternal(m, v940, int32(_a_F_ProcessInterrupts_27), int32(0), int32(_a_F_ProcessInterrupts_17), int32(_a_F_ProcessInterrupts_18))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L18
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	F_MemoryContextReset(m, v935)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L18
	} else {
		goto L273
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[50])) = v945
	v952 = v945
	goto L268
L273:
	;
	v951 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[50]))
	v952 = v951
	goto L268
L274:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v962)+4))
	if int32(0) < v963 {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	v1108 = v952
	goto L276
L276:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[42])) = v955
	F_MemoryContextReset(m, v1108)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L18
	} else {
		goto L315
	}
L277:
	;
	v968 = v953
	goto L280
L278:
	;
	goto L279
L279:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[50]))
	v1108 = v1106
	goto L276
L280:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v962)+12))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v977+v968<<(uint(int32(2))%32))))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v981)+4))
	if v982 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	goto L279
L282:
	;
	v1091 = v968 + int32(1)
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v962)+4))
	if v1091 < v1092 {
		v968 = v1091
		goto L280
	} else {
		goto L314
	}
L283:
	;
	v990 = F_shm_mq_receive(m, v982, v926+int32(72), v926+int32(68), int32(1))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L18
	} else {
		goto L287
	}
L284:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v926)+52))
	F_pfree(m, v1086)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L18
	} else {
		goto L313
	}
L285:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L18
	} else {
		goto L309
	}
L286:
	;
	v993 = v926 + int32(52)
	F_initStringInfo(m, v993)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L18
	} else {
		goto L288
	}
L287:
	;
	switch v990 {
	case 0:
		goto L286
	case 1:
		goto L282
	default:
		goto L285
	}
L288:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v926)+68))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v926)+72))
	F_appendBinaryStringInfo(m, v993, v996, v997)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L18
	} else {
		goto L289
	}
L289:
	;
	v1000 = F_pq_getmsgbyte(m, v993)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L18
	} else {
		goto L292
	}
L290:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L18
	} else {
		goto L306
	}
L291:
	;
	F_pq_parse_errornotice(m, v926+int32(52), v926+int32(76))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L18
	} else {
		goto L293
	}
L292:
	;
	v1002 = base.I32_extend8_s(v1000)
	switch v1002 - int32(65) {
	case 0, 13:
		goto L284
	default:
		goto L290
	case 4:
		goto L291
	}
L293:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v926)+124))
	if v1011 != 0 {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[52]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[46])) = v1026
	*(*int32)(unsafe.Add(mBase, uint32(v926)+124)) = v1023
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L18
	} else {
		goto L300
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v926)+36)) = int32(_a_F_ProcessInterrupts_28)
	*(*int32)(unsafe.Add(mBase, uint32(v926)+32)) = v1011
	v1018 = F_psprintf(m, int32(_a_F_ProcessInterrupts_23), v926+int32(32))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L18
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v1021 = F_pstrdup(m, int32(_a_F_ProcessInterrupts_28))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L18
	} else {
		goto L299
	}
L298:
	;
	v1023 = v1018
	goto L294
L299:
	;
	v1023 = v1021
	goto L294
L300:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L18
	} else {
		goto L301
	}
L301:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_29), int32(0))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L18
	} else {
		goto L302
	}
L302:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L18
	} else {
		goto L303
	}
L303:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v926)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v926)+16)) = v1043
	F_errcontext_msg(m, int32(_a_F_ProcessInterrupts_30), v926+int32(16))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L18
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_31), int32(1048), int32(_a_F_ProcessInterrupts_32))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L18
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v926))) = v1002
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v926)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v926)+4)) = v1060
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_33), v926)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L18
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_31), int32(1062), int32(_a_F_ProcessInterrupts_32))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L18
	} else {
		goto L308
	}
L308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L309:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L18
	} else {
		goto L310
	}
L310:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_34), int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L18
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_31), int32(1134), int32(_a_F_ProcessInterrupts_27))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L18
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	goto L282
L314:
	;
	goto L281
L315:
	;
	v1122 = int32(_a_F_ProcessInterrupts_13)
	v1124 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v1124 - int32(1)
	m.G0 = v926 + int32(176)
	goto L12
L316:
	;
	F_errcode(m, int32(50463042))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L18
	} else {
		goto L317
	}
L317:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_35), int32(0))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L18
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3593), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L18
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	F_errcode(m, int32(67240258))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L18
	} else {
		goto L321
	}
L321:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_36), int32(0))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L18
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3606), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L18
	} else {
		goto L323
	}
L323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L324:
	;
	F_errcode(m, int32(84017605))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L18
	} else {
		goto L325
	}
L325:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_37), int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L18
	} else {
		goto L326
	}
L326:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3619), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L18
	} else {
		goto L327
	}
L327:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L328:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L18
	} else {
		goto L329
	}
L329:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L18
	} else {
		goto L330
	}
L330:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_38), int32(0))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L18
	} else {
		goto L331
	}
L331:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3572), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L18
	} else {
		goto L332
	}
L332:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L333:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L18
	} else {
		goto L334
	}
L334:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L18
	} else {
		goto L335
	}
L335:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_39), int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L18
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3559), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L18
	} else {
		goto L337
	}
L337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L18
	} else {
		goto L339
	}
L339:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L18
	} else {
		goto L340
	}
L340:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_40), int32(0))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L18
	} else {
		goto L341
	}
L341:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3552), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L18
	} else {
		goto L342
	}
L342:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L343:
	;
	v1253 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[7])) = v1253
	F_errstart_cold(m, int32(22), v1253)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L18
	} else {
		goto L344
	}
L344:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L18
	} else {
		goto L345
	}
L345:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_41), int32(0))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L18
	} else {
		goto L346
	}
L346:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3493), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L18
	} else {
		goto L347
	}
L347:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L348:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L18
	} else {
		goto L349
	}
L349:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_42), int32(0))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L18
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3431), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L18
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L18
	} else {
		goto L353
	}
L353:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_43), int32(0))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L18
	} else {
		goto L354
	}
L354:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3427), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L18
	} else {
		goto L355
	}
L355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L356:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L18
	} else {
		goto L357
	}
L357:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L18
	} else {
		goto L358
	}
L358:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_44), int32(0))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L18
	} else {
		goto L359
	}
L359:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3545), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L18
	} else {
		goto L360
	}
L360:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L361:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L18
	} else {
		goto L362
	}
L362:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_45), int32(0))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L18
	} else {
		goto L363
	}
L363:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_1), int32(3423), int32(_a_F_ProcessInterrupts_2))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L18
	} else {
		goto L364
	}
L364:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessNotifyInterrupt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	goto L2
L1:
	;
	return
L2:
	;
	if v4 != int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[1]))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L5
L5:
	;
	v13 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[1])) = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[2]))
	if v16 == v13 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L1
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[1]))
	if v69 != 0 {
		goto L5
	} else {
		goto L28
	}
L8:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[3])))
	if v20 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L16
	}
L10:
	;
	v25 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	if v25 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessNotifyInterrupt_0), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_ProcessNotifyInterrupt_1), int32(2327), int32(_a_F_ProcessNotifyInterrupt_0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	F_asyncQueueReadAllNotifications(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	if l0 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[4]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v47 = m.T0[v46].(func(*base.Module) int32)(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[3])))
	if v50 != int32(1) {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v55 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	if v55 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessNotifyInterrupt_2), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_ProcessNotifyInterrupt_1), int32(2351), int32(_a_F_ProcessNotifyInterrupt_0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	goto L7
L28:
	;
	goto L6
}
func F_PromoteIsTriggered(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[0])))
	if v5 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
		v12 = base.AtomicRmwXchg32(m, v9, int32(96), int32(1))
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
			F_s_lock(m, v14+int32(96), int32(_a_F_PromoteIsTriggered_0), int32(_a_F_PromoteIsTriggered_1), int32(_a_F_PromoteIsTriggered_2))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
				*(*uint8)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[0])) = uint8(v27)
				v29 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+96)), uint32(v29))
				v32 = v27
				return v32 & int32(1)
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
			*(*uint8)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[0])) = uint8(v27)
			v29 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+96)), uint32(v29))
			v32 = v27
			return v32 & int32(1)
		}
	} else {
		v32 = int32(1)
		return v32 & int32(1)
	}
}
func F_p_isasclet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v6 != int32(1) {
		v74 = int32(0)
		return v74
	} else {
		v9 = int32(0)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v13 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10+v11))))
		if v13 < v9 {
			v74 = v9
			return v74
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
			if v16 == int32(1) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v19 != 0 {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20<<(uint(int32(2))%32))))
					return base.B2i32(base.Ui32(int32(127)) < base.Ui32(v24)) | base.B2i32(base.Ui32(v24|int32(32)-int32(97)) < base.Ui32(int32(26)))
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36<<(uint(int32(2))%32))))
					if base.Ui32(v40) <= base.Ui32(int32(_a_F_p_isasclet_0)) {
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v40)>>(uint(int32(8))%32)))+uint32(_c_F_p_isasclet[0]))))
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v40)>>(uint(int32(3))%32))&int32(31)|v49<<(uint(int32(5))%32))+uint32(_c_F_p_isasclet[0]))))
						v61 = int32(base.Ui32(v53)>>(uint(v40&int32(7))%32)) & int32(1)
					} else {
						v61 = base.B2i32(base.Ui32(v40) < base.Ui32(int32(_a_F_p_isasclet_1)))
					}
					return base.B2i32(v61 != int32(0))
				}
			} else {
				v74 = base.B2i32(base.Ui32((v13|int32(32)-int32(97))&int32(255)) < base.Ui32(int32(26)))
				return v74
			}
		}
	}
}
func F_p_isnotalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v67 int32
	_ = v67
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6+v9<<(uint(int32(2))%32))))
			if base.Ui32(int32(127)) < base.Ui32(v13) {
				v67 = int32(1)
				return base.B2i32(v67 == int32(0))
			} else {
				return base.B2i32(base.B2i32(base.Ui32(v13-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v13|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+v32<<(uint(int32(2))%32))))
			if base.Ui32(int32(10)) <= base.Ui32(v36-int32(48)) {
				v41 = F_iswalpha(m, v36)
				mBase = m.M
				v45 = base.B2i32(v41 != int32(0))
			} else {
				v45 = int32(1)
			}
			return base.B2i32(v45 == int32(0))
		}
	} else {
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
		v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v51))))
		v67 = base.B2i32(base.Ui32(v53-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v53|int32(32)-int32(97)) < base.Ui32(int32(26)))
		return base.B2i32(v67 == int32(0))
	}
}
func F_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int64
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(256)
	m.G0 = v9
	if l4&int32(_a_F_pad_0)|base.B2i32(l2 <= l3) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = l2 - l3
	v18 = int32(256)
	v20 = base.B2i32(base.Ui32(v17) < base.Ui32(v18))
	if base.Ui32(v17) < base.Ui32(v18) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v9 + int32(256)
	return
L4:
	;
	v21 = v17
	goto L6
L5:
	;
	v21 = v18
	goto L6
L6:
	;
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v20 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	goto L7
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v2)
	v28 = v9 + v21
	*(*uint8)(unsafe.Add(mBase, uint32(v28-int32(1)))) = uint8(v2)
	if base.Ui32(v21) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v28-int32(3)))) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v28-int32(2)))) = uint8(v2)
	if base.Ui32(v21) < base.Ui32(int32(7)) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+3)) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v28-int32(4)))) = uint8(v2)
	if base.Ui32(v21) < base.Ui32(int32(9)) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v53 = (int32(0) - v9) & int32(3)
	v54 = v9 + v53
	v58 = v2 & int32(255) * int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v58
	v62 = (v21 - v53) & int32(-4)
	v63 = v54 + v62
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(4)))) = v58
	if base.Ui32(v62) < base.Ui32(int32(9)) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(8)))) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(12)))) = v58
	if base.Ui32(v62) < base.Ui32(int32(25)) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(16)))) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(20)))) = v58
	v89 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v63-v89))) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(28)))) = v58
	v98 = v54&int32(4) | v89
	v99 = v62 - v98
	if base.Ui32(v99) < base.Ui32(int32(32)) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v104 = base.I64_extend_i32_u(v58) * int64(4294967297)
	v107 = v98 + v54
	v108 = v99
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v107)+24)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v107)+16)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v107)+8)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v107))) = v104
	v116 = int32(32)
	v119 = v108 - v116
	if base.Ui32(int32(31)) < base.Ui32(v119) {
		v107 = v107 + v116
		v108 = v119
		goto L16
	} else {
		goto L18
	}
L17:
	;
	goto L8
L18:
	;
	goto L17
L19:
	;
	v133 = v17
	goto L22
L20:
	;
	v146 = v17
	goto L21
L21:
	;
	F_out(m, l0, v9, v146)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L24
	} else {
		goto L27
	}
L22:
	;
	F_out(m, l0, v9, int32(256))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v146 = v140
	goto L21
L24:
	;
	return
L25:
	;
	v140 = v133 - int32(256)
	if base.Ui32(int32(255)) < base.Ui32(v140) {
		v133 = v140
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	goto L3
}
func F_pair_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v37 int32
	_ = v37
	var v41 float64
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = l0
	goto L1
L1:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v23-int32(9)))&base.B2i32(v23 != int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v37 = v14 + base.B2i32(v23 == int32(40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v37
	v41 = F_float8in_internal(m, v37, v12+int32(12), l4, l5, l6)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v14 = v14 + int32(1)
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	return int32(0)
L7:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v41
	if l6 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v12 + int32(16)
	return v146
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v57 = v55 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != int32(44) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v48 != int32(447) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v51 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v146 = int32(0)
	goto L8
L13:
	;
	v128 = int32(0)
	v129 = F_errsave_start(m, l6)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L30
	}
L14:
	;
	v64 = F_float8in_internal(m, v57, v12+int32(12), l4, l5, l6)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = v64
	if l6 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v23 != int32(40) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v69 != int32(447) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v72 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v146 = int32(0)
	goto L8
L20:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if l3 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v80 = v78 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 != int32(41) {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v85 = v80
	goto L23
L23:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v94-int32(9)))&base.B2i32(v94 != int32(32)) != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v103 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v103
	v85 = v103
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v114
	v146 = int32(1)
	goto L8
L27:
	;
	goto L28
L28:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v117 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v146 = int32(1)
	goto L8
L30:
	;
	if v129 == int32(0) {
		v146 = v128
		goto L8
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l4
	F_errmsg(m, int32(_a_F_pair_decode_0), v12)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	F_errsave_finish(m, l6, int32(_a_F_pair_decode_1), int32(251), int32(_a_F_pair_decode_2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v146 = v128
	goto L8
}
func F_pairingheap_remove(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = F_pairingheap_remove_first(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v18 != l1 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	v20 = int32(4)
	goto L8
L7:
	;
	v20 = int32(0)
	goto L8
L8:
	;
	v21 = v15 + v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v24 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
	if v22 == int32(0) {
		goto L9
	} else {
		goto L49
	}
L13:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v89
	if v22 == int32(0) {
		goto L9
	} else {
		goto L48
	}
L14:
	;
	v89 = v23
	goto L13
L15:
	;
	goto L16
L16:
	;
	v29 = int32(0)
	v30 = v23
	goto L17
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v37 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v29 == int32(0) {
		v89 = v57
		goto L13
	} else {
		goto L34
	}
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v29
	v57 = v30
	goto L19
L21:
	;
	goto L22
L22:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = m.T0[v43].(func(*base.Module, int32, int32, int32) int32)(m, v30, v37, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v47 = base.B2i32(v44 < int32(0))
	if v44 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v48 = v30
	goto L26
L25:
	;
	v48 = v37
	goto L26
L26:
	;
	if v44 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v49 = v37
	goto L29
L28:
	;
	v49 = v30
	goto L29
L29:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v48
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v49
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v48
	if v41 != 0 {
		v29 = v49
		v30 = v41
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v57 = v49
	goto L19
L34:
	;
	v64 = v57
	v66 = v29
	goto L35
L35:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = m.T0[v74].(func(*base.Module, int32, int32, int32) int32)(m, v64, v66, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	v89 = v80
	goto L13
L37:
	;
	v78 = base.B2i32(v75 < int32(0))
	if v75 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v79 = v64
	goto L40
L39:
	;
	v79 = v66
	goto L40
L40:
	;
	if v75 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v80 = v66
	goto L43
L42:
	;
	v80 = v64
	goto L43
L43:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v81 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v79
	goto L46
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v80
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v79
	if v72 != 0 {
		v64 = v80
		v66 = v72
		goto L35
	} else {
		goto L47
	}
L47:
	;
	goto L36
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v89
	return
L49:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v108
	goto L9
}
func F_paramlist_parser_setup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(816)
	return
}
func F_parse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v371 int32
	_ = v371
	v6 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = m.T0[v16].(func(*base.Module) int32)(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if v13 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v25 = v23
	goto L8
L7:
	;
	v25 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v25
	return int32(0)
L9:
	;
	v49 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v48)+4)) = v49
	v51 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)) = uint8(v51)
	v53 = int32(380)
	*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v48)+20)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(281479271677952)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v63 != 0 {
		v371 = v6
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v29
	v48 = v13
	goto L9
L11:
	;
	goto L12
L12:
	;
	v33 = F_palloc_extended(m, int32(88), int32(2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v33 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v39 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v33
	v48 = v33
	goto L9
L17:
	;
	v41 = v39
	goto L19
L18:
	;
	v41 = int32(12)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v41
	return int32(0)
L20:
	;
	return v371
L21:
	;
	v65 = v48 + int32(20)
	v77 = v6
	goto L23
L22:
	;
	if l1 != v278 {
		goto L94
	} else {
		goto L95
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v79 = F_newstate(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v278 = v277
	goto L22
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v82 = F_newstate(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_parse[0]))
	if v89 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v92 <= v93 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L32
L34:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_parse[0]))
	if v169 != 0 {
		goto L56
	} else {
		goto L57
	}
L35:
	;
	F_createarc(m, v87, int32(110), int32(0), l3, v79)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L55
	}
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v95 == int32(0) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	if v117 == int32(0) {
		goto L35
	} else {
		goto L47
	}
L39:
	;
	v103 = v95
	goto L40
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	if v110 != v79 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L35
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	if v116 != 0 {
		v103 = v116
		goto L40
	} else {
		goto L46
	}
L43:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
	if v112 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v113 == int32(110) {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	goto L41
L47:
	;
	v125 = v117
	goto L48
L48:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	if v132 != l3 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L35
L50:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v125)+24))
	if v138 != 0 {
		v125 = v138
		goto L48
	} else {
		goto L54
	}
L51:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+4)))
	if v134 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v135 == int32(110) {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	goto L49
L55:
	;
	goto L34
L56:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v172 <= v173 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L58
L60:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v247 != 0 {
		goto L82
	} else {
		goto L83
	}
L61:
	;
	F_createarc(m, v167, int32(110), int32(0), v82, l4)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L81
	}
L62:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v175 == int32(0) {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v197 == int32(0) {
		goto L61
	} else {
		goto L73
	}
L65:
	;
	v183 = v175
	goto L66
L66:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	if v190 != l4 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L61
L68:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	if v196 != 0 {
		v183 = v196
		goto L66
	} else {
		goto L72
	}
L69:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+4)))
	if v192 != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v193 == int32(110) {
		goto L60
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	goto L67
L73:
	;
	v205 = v197
	goto L74
L74:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	if v212 != v82 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L61
L76:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v205)+24))
	if v218 != 0 {
		v205 = v218
		goto L74
	} else {
		goto L80
	}
L77:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+4)))
	if v214 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if v215 == int32(110) {
		goto L60
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	goto L75
L81:
	;
	goto L60
L82:
	;
	return int32(0)
L83:
	;
	goto L84
L84:
	;
	v250 = int32(0)
	v252 = F_parsebranch(m, l0, l1, l2, v79, v82, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v254 != 0 {
		v371 = v250
		goto L20
	} else {
		goto L86
	}
L86:
	;
	if v77 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+1)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v261 = v257 | v260
	v270 = v257&int32(28) | v261<<(uint(int32(1))%32)&(v261<<(uint(int32(2))%32))&int32(4) | v260
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v270)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v272 != int32(124) {
		v278 = v272
		goto L22
	} else {
		goto L91
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = v252
	goto L87
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v252
	goto L87
L91:
	;
	v275 = F_next(m, l0)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v275 != 0 {
		v77 = v252
		goto L23
	} else {
		goto L93
	}
L93:
	;
	goto L24
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v283 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v288 == v252 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v285 = v283
	goto L99
L98:
	;
	v285 = int32(8)
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v285
	goto L96
L100:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
	if v290 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v316&int32(28) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L103:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v48)+64))
	F_pfree(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v302 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v302)
	v304 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v304
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v304
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v308 != 0 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	F_pfree(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
	F_pfree(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = int32(0)
	goto L105
L109:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v48
	return v252
L110:
	;
	goto L111
L111:
	;
	F_pfree(m, v48)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	return v252
L113:
	;
	if v288 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	v371 = v48
	goto L20
L116:
	;
	v328 = v288
	goto L119
L117:
	;
	goto L118
L118:
	;
	v348 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v348)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(0)
	goto L115
L119:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	F_freesubre(m, l0, v328)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	goto L118
L121:
	;
	if v333 != 0 {
		v328 = v333
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
}
func F_parseCheckAggregates(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v686 int32
	_ = v686
	var v687 int64
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int64
	_ = v731
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v3
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v25 == v3 {
		v255 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v262 == int32(0) {
		v369 = v3
		v374 = v3
		goto L41
	} else {
		goto L42
	}
L2:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+104)))
	v30 = F_expand_grouping_sets(m, v25, v28, int32(_a_F_parseCheckAggregates_0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v30 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v57 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	F_errcode(m, int32(16777477))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(_a_F_parseCheckAggregates_1), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = v45
	goto L13
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v47 = v46
	goto L13
L13:
	;
	v48 = F_exprLocation(m, v47)
	mBase = m.M
	F_parser_errposition(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_parseCheckAggregates_2), int32(1172), int32(_a_F_parseCheckAggregates_3))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v236 != int32(1) {
		v255 = v229
		goto L1
	} else {
		goto L39
	}
L17:
	;
	v229 = int32(0)
	goto L16
L18:
	;
	v60 = int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v61 <= v60 {
		v229 = v57
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v68 = v60
	v75 = v57
	goto L20
L20:
	;
	v82 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v68<<(uint(int32(2))%32))))
	if base.B2i32(v75 == v82)|base.B2i32(v90 == v82) != 0 {
		v178 = v82
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v229 = v178
	goto L16
L22:
	;
	if v178 == int32(0) {
		goto L17
	} else {
		goto L37
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v94 <= int32(0) {
		v178 = v82
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v100 = v82
	v105 = v82
	v106 = v94
	goto L25
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v115 <= int32(0) {
		v157 = v100
		v163 = v106
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v178 = v157
	goto L22
L27:
	;
	v173 = v105 + int32(1)
	if v173 < v163 {
		v100 = v157
		v105 = v173
		v106 = v163
		goto L25
	} else {
		goto L36
	}
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+v105<<(uint(int32(2))%32))))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v130 = int32(0)
	goto L29
L29:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v123+v130<<(uint(int32(2))%32))))
	if v122 != v146 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v151 = F_lappend_int(m, v100, v122)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L35
	}
L31:
	;
	v149 = v130 + int32(1)
	if v149 != v115 {
		v130 = v149
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	v157 = v100
	v163 = v106
	goto L27
L35:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v157 = v151
	v163 = v153
	goto L27
L36:
	;
	goto L26
L37:
	;
	v196 = v68 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v196 < v197 {
		v68 = v196
		v75 = v178
		goto L20
	} else {
		goto L38
	}
L38:
	;
	goto L21
L39:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v239 == int32(0) {
		v255 = v229
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = int32(0)
	v255 = v229
	goto L1
L41:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v377 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L42:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v265 <= int32(0) {
		v369 = v3
		v374 = v3
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v268 = int32(0)
	if v265 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v272 = int32(0)
	if v272 < v265 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v332 = v268
	v333 = v268
	v344 = v3
	goto L46
L46:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v349 = int32(2)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348+v333<<(uint(v349)%32))))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+12))
	switch v353 - v349 {
	case 0:
		v369 = int32(1)
		v374 = v344
		goto L41
	default:
		v358 = v344
		goto L60
	case 4:
		goto L61
	}
L47:
	;
	v275 = v265
	goto L49
L48:
	;
	v275 = v272
	goto L49
L49:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v285 = v268
	v287 = v3
	v295 = v3
	v296 = v3
	goto L50
L50:
	;
	v299 = int32(1)
	v301 = int32(2)
	v303 = v280 + v285<<(uint(v301)%32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	switch v305 - v301 {
	case 0:
		v311 = v299
		v312 = v296
		goto L52
	default:
		v310 = v296
		goto L53
	case 4:
		goto L54
	}
L51:
	;
	if v275&int32(1) == int32(0) {
		v369 = v320
		v374 = v321
		goto L41
	} else {
		goto L59
	}
L52:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+12))
	switch v314 - int32(2) {
	case 0:
		v320 = v299
		v321 = v312
		goto L55
	default:
		v319 = v312
		goto L56
	case 4:
		goto L57
	}
L53:
	;
	v311 = v287
	v312 = v310
	goto L52
L54:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+92)))
	v310 = v308 | v296
	goto L53
L55:
	;
	v322 = int32(2)
	v323 = v285 + v322
	v325 = v295 + v322
	if v325 != v275&int32(2147483646) {
		v285 = v323
		v287 = v320
		v295 = v325
		v296 = v321
		goto L50
	} else {
		goto L58
	}
L56:
	;
	v320 = v311
	v321 = v319
	goto L55
L57:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+92)))
	v319 = v317 | v312
	goto L56
L58:
	;
	goto L51
L59:
	;
	v332 = v320
	v333 = v323
	v344 = v321
	goto L46
L60:
	;
	v369 = v332
	v374 = v358
	goto L41
L61:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+92)))
	v358 = v356 | v344
	goto L60
L62:
	;
	if v369&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L63:
	;
	v423 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v381 = int32(0)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	if v382 <= v381 {
		v423 = v381
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v389 = int32(0)
	v390 = v381
	goto L67
L67:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v377)+12))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v404+v389<<(uint(int32(2))%32))))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v410 = F_get_sortgroupclause_tle(m, v408, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L3
	} else {
		goto L69
	}
L68:
	;
	v423 = v414
	goto L62
L69:
	;
	if v410 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v412 = F_lappend(m, v390, v410)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L3
	} else {
		goto L73
	}
L71:
	;
	v414 = v390
	goto L72
L72:
	;
	v416 = v389 + int32(1)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	if v416 < v417 {
		v389 = v416
		v390 = v414
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v414 = v412
	goto L72
L74:
	;
	goto L68
L75:
	;
	v440 = F_flatten_join_alias_vars(m, int32(0), l1, v423)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L3
	} else {
		goto L78
	}
L76:
	;
	v442 = v423
	goto L77
L77:
	;
	if v442 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v442 = v440
	goto L77
L79:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v687 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+36)) = v687
	v689 = int32(1)
	v690 = v674 & v689
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v690)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v687
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v442
	v696 = v369 & v689
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)) = uint8(v696)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	v700 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v700)
	v704 = F_finalize_grouping_exprs_walker(m, v686, v21+int32(8))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L3
	} else {
		goto L139
	}
L80:
	;
	v445 = int32(0)
	v673 = v445
	v674 = v445
	goto L79
L81:
	;
	goto L82
L82:
	;
	v447 = int32(0)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	if v448 <= v447 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v553 = int32(0)
	v559 = F_palloc0(m, int32(136))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L3
	} else {
		goto L112
	}
L84:
	;
	v540 = int32(0)
	v541 = v447
	goto L83
L85:
	;
	goto L86
L86:
	;
	v452 = int32(0)
	v457 = v452
	v459 = v452
	v460 = v447
	goto L87
L87:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v442)+12))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v472+v457<<(uint(int32(2))%32))))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	if v478 != int32(6) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v540 = v529
	v541 = v530
	goto L83
L89:
	;
	v532 = v457 + int32(1)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	if v532 < v533 {
		v457 = v532
		v459 = v529
		v460 = v530
		goto L87
	} else {
		goto L111
	}
L90:
	;
	v529 = v459
	v530 = int32(1)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v482 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v476)+16))
	v484 = int32(0)
	if v255 == v484 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v526 = v477
	goto L95
L95:
	;
	v527 = F_lappend(m, v459, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L3
	} else {
		goto L110
	}
L96:
	;
	if v522 == int32(0) {
		v529 = v459
		v530 = v460
		goto L89
	} else {
		goto L109
	}
L97:
	;
	v522 = int32(0)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v490 <= int32(0) {
		v516 = v484
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v522 = v516
	goto L96
L101:
	;
	v493 = int32(0)
	if v493 < v490 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v496 = v490
	goto L104
L103:
	;
	v496 = v493
	goto L104
L104:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v499 = int32(0)
	goto L105
L105:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v497+v499<<(uint(int32(2))%32))))
	v508 = base.B2i32(v507 == v483)
	if v507 == v483 {
		v516 = v508
		goto L100
	} else {
		goto L107
	}
L106:
	;
	v516 = v508
	goto L100
L107:
	;
	v510 = v499 + int32(1)
	if v510 != v496 {
		v499 = v510
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	v526 = v525
	goto L95
L110:
	;
	v529 = v527
	v530 = v460
	goto L89
L111:
	;
	goto L88
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559)+12)) = int32(9)
	*(*int64)(unsafe.Add(mBase, uint32(v559))) = int64(101)
	v567 = F_makeAlias(m, int32(_a_F_parseCheckAggregates_4), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	if v442 == int32(0) {
		v640 = v553
		v641 = v553
		v644 = v553
		v645 = v553
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v650 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v559)+124)) = uint16(v650)
	*(*int32)(unsafe.Add(mBase, uint32(v559)+120)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(v559)+8)) = v567
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v655 = F_lappend(m, v654, v559)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L3
	} else {
		goto L134
	}
L115:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	if v571 <= int32(0) {
		v640 = v553
		v641 = v553
		v644 = v553
		v645 = v553
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v582 = v553
	v583 = v553
	v586 = v553
	v587 = v553
	v590 = v553
	goto L117
L117:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v442)+12))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v592+v590<<(uint(int32(2))%32))))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	if v597 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v640 = v621
	v641 = v616
	v644 = v611
	v645 = v626
	goto L114
L119:
	;
	v598 = F_pstrdup(m, v597)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L3
	} else {
		goto L122
	}
L120:
	;
	v601 = int32(_a_F_parseCheckAggregates_5)
	goto L121
L121:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	v603 = F_makeString(m, v601)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L3
	} else {
		goto L123
	}
L122:
	;
	v601 = v598
	goto L121
L123:
	;
	v605 = F_lappend(m, v602, v603)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567)+8)) = v605
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	v609 = F_copyObjectImpl(m, v608)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	v611 = F_lappend(m, v586, v609)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	v614 = F_exprType(m, v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v616 = F_lappend_oid(m, v583, v614)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	v619 = F_exprTypmod(m, v618)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	v621 = F_lappend_int(m, v582, v619)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	v624 = F_exprCollation(m, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	v626 = F_lappend_oid(m, v587, v624)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	v629 = v590 + int32(1)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	if v629 < v630 {
		v582 = v621
		v583 = v616
		v586 = v611
		v587 = v626
		v590 = v629
		goto L117
	} else {
		goto L133
	}
L133:
	;
	goto L118
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v655
	if v655 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v655)+4))
	v660 = v658
	goto L137
L136:
	;
	v660 = int32(0)
	goto L137
L137:
	;
	v661 = F_buildNSItemFromLists(m, v559, v660, v641, v640, v645)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L3
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v661
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v665 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v665)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v664
	v673 = v540
	v674 = v541
	goto L79
L139:
	;
	if v696 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v707 = F_flatten_join_alias_vars(m, int32(0), l1, v686)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L3
	} else {
		goto L143
	}
L141:
	;
	v709 = v686
	goto L142
L142:
	;
	v710 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v710)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v710
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v690)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)) = uint8(v710)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(4)
	v726 = v21 + int32(8)
	v727 = F_substitute_grouped_columns_mutator(m, v709, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L3
	} else {
		goto L144
	}
L143:
	;
	v709 = v707
	goto L142
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v727
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v731 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+36)) = v731
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v690)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)) = uint8(v696)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	v740 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v740)
	v742 = F_finalize_grouping_exprs_walker(m, v730, v726)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L3
	} else {
		goto L145
	}
L145:
	;
	if v696 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v745 = F_flatten_join_alias_vars(m, int32(0), l1, v730)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L3
	} else {
		goto L149
	}
L147:
	;
	v747 = v730
	goto L148
L148:
	;
	v748 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v748)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v748
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v690)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+16)) = uint8(v748)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	v765 = F_substitute_grouped_columns_mutator(m, v747, v21+int32(8))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L3
	} else {
		goto L150
	}
L149:
	;
	v747 = v745
	goto L148
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v765
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v768&v374&int32(1) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L3
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	m.G0 = v21 + int32(48)
	return
L154:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L3
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(_a_F_parseCheckAggregates_6), int32(0))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L3
	} else {
		goto L156
	}
L156:
	;
	v784 = F_locate_agg_of_level(m, l1, int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	F_parser_errposition(m, l0, v784)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L3
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_parseCheckAggregates_2), int32(1331), int32(_a_F_parseCheckAggregates_3))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L3
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_datetime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v114 int32
	_ = v114
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int64
	_ = v226
	var v227 int64
	_ = v227
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int64
	_ = v421
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v426 int64
	_ = v426
	var v427 int64
	_ = v427
	var v429 int64
	_ = v429
	var v430 int64
	_ = v430
	var v433 int64
	_ = v433
	var v442 int32
	_ = v442
	var v443 int64
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v468 int64
	_ = v468
	var v469 int64
	_ = v469
	var v470 int64
	_ = v470
	var v473 int64
	_ = v473
	var v474 int64
	_ = v474
	var v476 int64
	_ = v476
	var v477 int64
	_ = v477
	var v480 int64
	_ = v480
	var v488 int64
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v570 int32
	_ = v570
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v17 = v12 + int32(52)
	v26 = F_do_to_timestamp(m, l0, l1, int32(100), int32(1), v17, v12+int32(40), v12+int32(44), v12+int32(36), v12+int32(32), l5)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		if v26 == int32(0) {
			v570 = int32(0)
			m.G0 = v12 + int32(96)
			return v570
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
			if v32 != 0 {
				v34 = v32
			} else {
				v34 = int32(-1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v34
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
			v38 = v36 & int32(2)
			if v36&int32(1) != 0 {
				v42 = v36 & int32(4)
				if v38 != 0 {
					if v42 != 0 {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
						if v43 == int32(1) {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
							v50 = v12 + int32(24)
							v57 = m.G0
							v59 = v57 - int32(16)
							m.G0 = v59
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
							if v61 <= int32(-4713) {
								if v61 != int32(-4713) {
									*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
									v138 = int32(-1)
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
									if int32(10) < v66 {
										v77 = v66
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
										v79 = F_date2j(m, v61, v77, v78)
										mBase = m.M
										v82 = base.I64_extend_i32_s(v79 - int32(_a_F_parse_datetime_0))
										v83 = int64(63)
										F___multi3(m, v59, v82, v82>>(uint(v83)%64), int64(86400000000), int64(0))
										mBase = m.M
										v88 = *(*int64)(unsafe.Add(mBase, uint32(v59)+8))
										v89 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
										if v88 != v89>>(uint(v83)%64) {
											*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
											v138 = int32(-1)
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
											v97 = int32(60)
											v106 = base.I64_extend_i32_s(v48) + base.I64_extend_i32_s(v94+(v95+v96*v97)*v97)*int64(1000000)
											v107 = v89 + v106
											*(*int64)(unsafe.Add(mBase, uint32(v50))) = v107
											if base.B2i32(v106 < int64(0))^base.B2i32(v107 < v89) != 0 {
												*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
												v138 = int32(-1)
											} else {
												if l4 != 0 {
													v114 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
													v119 = base.I64_extend_i32_s(int32(0)-v114)*int64(-1000000) + v107
													*(*int64)(unsafe.Add(mBase, uint32(v50))) = v119
													v121 = v119
												} else {
													v121 = v107
												}
												if base.Ui64(v121+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
													v138 = int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
													v138 = int32(-1)
												}
											}
										}
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
										v138 = int32(-1)
									}
								}
							} else {
								if v61 <= int32(_a_F_parse_datetime_1) {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
									v77 = v71
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
									v79 = F_date2j(m, v61, v77, v78)
									mBase = m.M
									v82 = base.I64_extend_i32_s(v79 - int32(_a_F_parse_datetime_0))
									v83 = int64(63)
									F___multi3(m, v59, v82, v82>>(uint(v83)%64), int64(86400000000), int64(0))
									mBase = m.M
									v88 = *(*int64)(unsafe.Add(mBase, uint32(v59)+8))
									v89 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
									if v88 != v89>>(uint(v83)%64) {
										*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
										v138 = int32(-1)
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
										v97 = int32(60)
										v106 = base.I64_extend_i32_s(v48) + base.I64_extend_i32_s(v94+(v95+v96*v97)*v97)*int64(1000000)
										v107 = v89 + v106
										*(*int64)(unsafe.Add(mBase, uint32(v50))) = v107
										if base.B2i32(v106 < int64(0))^base.B2i32(v107 < v89) != 0 {
											*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
											v138 = int32(-1)
										} else {
											if l4 != 0 {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
												v119 = base.I64_extend_i32_s(int32(0)-v114)*int64(-1000000) + v107
												*(*int64)(unsafe.Add(mBase, uint32(v50))) = v119
												v121 = v119
											} else {
												v121 = v107
											}
											if base.Ui64(v121+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v138 = int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
												v138 = int32(-1)
											}
										}
									}
								} else {
									if v61 != int32(_a_F_parse_datetime_2) {
										*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
										v138 = int32(-1)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
										if int32(5) < v74 {
											*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
											v138 = int32(-1)
										} else {
											v77 = v74
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
											v79 = F_date2j(m, v61, v77, v78)
											mBase = m.M
											v82 = base.I64_extend_i32_s(v79 - int32(_a_F_parse_datetime_0))
											v83 = int64(63)
											F___multi3(m, v59, v82, v82>>(uint(v83)%64), int64(86400000000), int64(0))
											mBase = m.M
											v88 = *(*int64)(unsafe.Add(mBase, uint32(v59)+8))
											v89 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
											if v88 != v89>>(uint(v83)%64) {
												*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
												v138 = int32(-1)
											} else {
												v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
												v97 = int32(60)
												v106 = base.I64_extend_i32_s(v48) + base.I64_extend_i32_s(v94+(v95+v96*v97)*v97)*int64(1000000)
												v107 = v89 + v106
												*(*int64)(unsafe.Add(mBase, uint32(v50))) = v107
												if base.B2i32(v106 < int64(0))^base.B2i32(v107 < v89) != 0 {
													*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
													v138 = int32(-1)
												} else {
													if l4 != 0 {
														v114 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
														v119 = base.I64_extend_i32_s(int32(0)-v114)*int64(-1000000) + v107
														*(*int64)(unsafe.Add(mBase, uint32(v50))) = v119
														v121 = v119
													} else {
														v121 = v107
													}
													if base.Ui64(v121+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
														v138 = int32(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v50))) = int64(0)
														v138 = int32(-1)
													}
												}
											}
										}
									}
								}
							}
							m.G0 = v59 + int32(16)
							if v138 == int32(0) {
								v180 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								F_AdjustTimestampForTypmod(m, v12+int32(24), v180, l5)
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1184)
									v185 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
									v186 = F_Int64GetDatum(m, v185)
									mBase = m.M
									v187 = m.ExcPending
									if v187 != 0 {
										return int32(0)
									} else {
										v570 = v186
										m.G0 = v12 + int32(96)
										return v570
									}
								}
							} else {
								v144 = int32(0)
								v145 = F_errsave_start(m, l5)
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									if v145 == int32(0) {
										v570 = v144
										m.G0 = v12 + int32(96)
										return v570
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_parse_datetime_3), int32(0))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_5), int32(_a_F_parse_datetime_6))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return int32(0)
												} else {
													v570 = v144
													m.G0 = v12 + int32(96)
													return v570
												}
											}
										}
									}
								}
							}
						} else {
							v161 = int32(0)
							v162 = F_errsave_start(m, l5)
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int32(0)
							} else {
								if v162 == int32(0) {
									v570 = v161
									m.G0 = v12 + int32(96)
									return v570
								} else {
									F_errcode(m, int32(117440642))
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_parse_datetime_7), int32(0))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_8), int32(_a_F_parse_datetime_6))
											mBase = m.M
											v177 = m.ExcPending
											if v177 != 0 {
												return int32(0)
											} else {
												v570 = v161
												m.G0 = v12 + int32(96)
												return v570
											}
										}
									}
								}
							}
						}
					} else {
						v188 = int32(0)
						v190 = v12 + int32(52)
						v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
						v194 = v12 + int32(24)
						v201 = m.G0
						v203 = v201 - int32(16)
						m.G0 = v203
						v205 = *(*int32)(unsafe.Add(mBase, uint32(v190)+20))
						if v205 <= int32(-4713) {
							if v205 != int32(-4713) {
								*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
								v282 = int32(-1)
							} else {
								v210 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
								if int32(10) < v210 {
									v221 = v210
									v222 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
									v223 = F_date2j(m, v205, v221, v222)
									mBase = m.M
									v226 = base.I64_extend_i32_s(v223 - int32(_a_F_parse_datetime_0))
									v227 = int64(63)
									F___multi3(m, v203, v226, v226>>(uint(v227)%64), int64(86400000000), int64(0))
									mBase = m.M
									v232 = *(*int64)(unsafe.Add(mBase, uint32(v203)+8))
									v233 = *(*int64)(unsafe.Add(mBase, uint32(v203)))
									if v232 != v233>>(uint(v227)%64) {
										*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
										v282 = int32(-1)
									} else {
										v238 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
										v239 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
										v240 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
										v241 = int32(60)
										v250 = base.I64_extend_i32_s(v191) + base.I64_extend_i32_s(v238+(v239+v240*v241)*v241)*int64(1000000)
										v251 = v233 + v250
										*(*int64)(unsafe.Add(mBase, uint32(v194))) = v251
										if base.B2i32(v250 < int64(0))^base.B2i32(v251 < v233) != 0 {
											*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
											v282 = int32(-1)
										} else {
											if base.Ui64(v251+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v282 = int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
												v282 = int32(-1)
											}
										}
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
									v282 = int32(-1)
								}
							}
						} else {
							if v205 <= int32(_a_F_parse_datetime_1) {
								v215 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
								v221 = v215
								v222 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
								v223 = F_date2j(m, v205, v221, v222)
								mBase = m.M
								v226 = base.I64_extend_i32_s(v223 - int32(_a_F_parse_datetime_0))
								v227 = int64(63)
								F___multi3(m, v203, v226, v226>>(uint(v227)%64), int64(86400000000), int64(0))
								mBase = m.M
								v232 = *(*int64)(unsafe.Add(mBase, uint32(v203)+8))
								v233 = *(*int64)(unsafe.Add(mBase, uint32(v203)))
								if v232 != v233>>(uint(v227)%64) {
									*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
									v282 = int32(-1)
								} else {
									v238 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
									v239 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
									v240 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
									v241 = int32(60)
									v250 = base.I64_extend_i32_s(v191) + base.I64_extend_i32_s(v238+(v239+v240*v241)*v241)*int64(1000000)
									v251 = v233 + v250
									*(*int64)(unsafe.Add(mBase, uint32(v194))) = v251
									if base.B2i32(v250 < int64(0))^base.B2i32(v251 < v233) != 0 {
										*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
										v282 = int32(-1)
									} else {
										if base.Ui64(v251+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
											v282 = int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
											v282 = int32(-1)
										}
									}
								}
							} else {
								if v205 != int32(_a_F_parse_datetime_2) {
									*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
									v282 = int32(-1)
								} else {
									v218 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
									if int32(5) < v218 {
										*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
										v282 = int32(-1)
									} else {
										v221 = v218
										v222 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
										v223 = F_date2j(m, v205, v221, v222)
										mBase = m.M
										v226 = base.I64_extend_i32_s(v223 - int32(_a_F_parse_datetime_0))
										v227 = int64(63)
										F___multi3(m, v203, v226, v226>>(uint(v227)%64), int64(86400000000), int64(0))
										mBase = m.M
										v232 = *(*int64)(unsafe.Add(mBase, uint32(v203)+8))
										v233 = *(*int64)(unsafe.Add(mBase, uint32(v203)))
										if v232 != v233>>(uint(v227)%64) {
											*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
											v282 = int32(-1)
										} else {
											v238 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
											v239 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
											v240 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
											v241 = int32(60)
											v250 = base.I64_extend_i32_s(v191) + base.I64_extend_i32_s(v238+(v239+v240*v241)*v241)*int64(1000000)
											v251 = v233 + v250
											*(*int64)(unsafe.Add(mBase, uint32(v194))) = v251
											if base.B2i32(v250 < int64(0))^base.B2i32(v251 < v233) != 0 {
												*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
												v282 = int32(-1)
											} else {
												if base.Ui64(v251+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
													v282 = int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v194))) = int64(0)
													v282 = int32(-1)
												}
											}
										}
									}
								}
							}
						}
						m.G0 = v203 + int32(16)
						if v282 != 0 {
							v286 = F_errsave_start(m, l5)
							mBase = m.M
							v287 = m.ExcPending
							if v287 != 0 {
								return int32(0)
							} else {
								if v286 == int32(0) {
									v570 = v188
									m.G0 = v12 + int32(96)
									return v570
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v292 = m.ExcPending
									if v292 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_parse_datetime_9), int32(0))
										mBase = m.M
										v296 = m.ExcPending
										if v296 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_10), int32(_a_F_parse_datetime_6))
											mBase = m.M
											v301 = m.ExcPending
											if v301 != 0 {
												return int32(0)
											} else {
												v570 = v188
												m.G0 = v12 + int32(96)
												return v570
											}
										}
									}
								}
							}
						} else {
							v304 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							F_AdjustTimestampForTypmod(m, v12+int32(24), v304, l5)
							mBase = m.M
							v306 = m.ExcPending
							if v306 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1114)
								v309 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
								v310 = F_Int64GetDatum(m, v309)
								mBase = m.M
								v311 = m.ExcPending
								if v311 != 0 {
									return int32(0)
								} else {
									v570 = v310
									m.G0 = v12 + int32(96)
									return v570
								}
							}
						}
					}
				} else {
					if v42 != 0 {
						v312 = int32(0)
						v313 = F_errsave_start(m, l5)
						mBase = m.M
						v314 = m.ExcPending
						if v314 != 0 {
							return int32(0)
						} else {
							if v313 == int32(0) {
								v570 = v312
								m.G0 = v12 + int32(96)
								return v570
							} else {
								F_errcode(m, int32(117440642))
								mBase = m.M
								v319 = m.ExcPending
								if v319 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_parse_datetime_11), int32(0))
									mBase = m.M
									v323 = m.ExcPending
									if v323 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_12), int32(_a_F_parse_datetime_6))
										mBase = m.M
										v328 = m.ExcPending
										if v328 != 0 {
											return int32(0)
										} else {
											v570 = v312
											m.G0 = v12 + int32(96)
											return v570
										}
									}
								}
							}
						}
					} else {
						v329 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
						if v329 <= int32(-4713) {
							if v329 != int32(-4713) {
								v346 = int32(0)
								v347 = F_errsave_start(m, l5)
								mBase = m.M
								v348 = m.ExcPending
								if v348 != 0 {
									return int32(0)
								} else {
									if v347 == int32(0) {
										v570 = v346
										m.G0 = v12 + int32(96)
										return v570
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v353 = m.ExcPending
										if v353 != 0 {
											return int32(0)
										} else {
											v354 = F_text_to_cstring(m, l0)
											mBase = m.M
											v355 = m.ExcPending
											if v355 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v354
												F_errmsg(m, int32(_a_F_parse_datetime_13), v12+int32(16))
												mBase = m.M
												v361 = m.ExcPending
												if v361 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
													mBase = m.M
													v366 = m.ExcPending
													if v366 != 0 {
														return int32(0)
													} else {
														v570 = v346
														m.G0 = v12 + int32(96)
														return v570
													}
												}
											}
										}
									}
								}
							} else {
								v334 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
								if v334 <= int32(10) {
									v346 = int32(0)
									v347 = F_errsave_start(m, l5)
									mBase = m.M
									v348 = m.ExcPending
									if v348 != 0 {
										return int32(0)
									} else {
										if v347 == int32(0) {
											v570 = v346
											m.G0 = v12 + int32(96)
											return v570
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v353 = m.ExcPending
											if v353 != 0 {
												return int32(0)
											} else {
												v354 = F_text_to_cstring(m, l0)
												mBase = m.M
												v355 = m.ExcPending
												if v355 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v354
													F_errmsg(m, int32(_a_F_parse_datetime_13), v12+int32(16))
													mBase = m.M
													v361 = m.ExcPending
													if v361 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
														mBase = m.M
														v366 = m.ExcPending
														if v366 != 0 {
															return int32(0)
														} else {
															v570 = v346
															m.G0 = v12 + int32(96)
															return v570
														}
													}
												}
											}
										}
									}
								} else {
									v508 = v334
									v509 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
									v514 = base.B2i32(int32(2) < v508)
									if int32(2) < v508 {
										v515 = int32(_a_F_parse_datetime_15)
									} else {
										v515 = int32(_a_F_parse_datetime_16)
									}
									v516 = v515 + v329
									v521 = base.I32_div_s(v516, int32(4))
									v524 = base.I32_div_s(v516, int32(-100))
									v527 = base.I32_div_s(v516, int32(400))
									if int32(2) < v508 {
										v531 = int32(1)
									} else {
										v531 = int32(13)
									}
									v536 = base.I32_div_s((v531+v508)*int32(_a_F_parse_datetime_17), int32(256))
									v539 = v509 + v516*int32(365) + v521 + v524 + v527 + v536 - int32(_a_F_parse_datetime_18)
									if base.Ui32(int32(2147483494)) <= base.Ui32(v539) {
										v542 = int32(0)
										v543 = F_errsave_start(m, l5)
										mBase = m.M
										v544 = m.ExcPending
										if v544 != 0 {
											return int32(0)
										} else {
											if v543 == int32(0) {
												v570 = v542
												m.G0 = v12 + int32(96)
												return v570
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v549 = m.ExcPending
												if v549 != 0 {
													return int32(0)
												} else {
													v550 = F_text_to_cstring(m, l0)
													mBase = m.M
													v551 = m.ExcPending
													if v551 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12))) = v550
														F_errmsg(m, int32(_a_F_parse_datetime_13), v12)
														mBase = m.M
														v555 = m.ExcPending
														if v555 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_19), int32(_a_F_parse_datetime_6))
															mBase = m.M
															v560 = m.ExcPending
															if v560 != 0 {
																return int32(0)
															} else {
																v570 = v542
																m.G0 = v12 + int32(96)
																return v570
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1082)
										v570 = v539 - int32(_a_F_parse_datetime_0)
										m.G0 = v12 + int32(96)
										return v570
									}
								}
							}
						} else {
							if v329 <= int32(_a_F_parse_datetime_1) {
								v339 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
								v508 = v339
								v509 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
								v514 = base.B2i32(int32(2) < v508)
								if int32(2) < v508 {
									v515 = int32(_a_F_parse_datetime_15)
								} else {
									v515 = int32(_a_F_parse_datetime_16)
								}
								v516 = v515 + v329
								v521 = base.I32_div_s(v516, int32(4))
								v524 = base.I32_div_s(v516, int32(-100))
								v527 = base.I32_div_s(v516, int32(400))
								if int32(2) < v508 {
									v531 = int32(1)
								} else {
									v531 = int32(13)
								}
								v536 = base.I32_div_s((v531+v508)*int32(_a_F_parse_datetime_17), int32(256))
								v539 = v509 + v516*int32(365) + v521 + v524 + v527 + v536 - int32(_a_F_parse_datetime_18)
								if base.Ui32(int32(2147483494)) <= base.Ui32(v539) {
									v542 = int32(0)
									v543 = F_errsave_start(m, l5)
									mBase = m.M
									v544 = m.ExcPending
									if v544 != 0 {
										return int32(0)
									} else {
										if v543 == int32(0) {
											v570 = v542
											m.G0 = v12 + int32(96)
											return v570
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v549 = m.ExcPending
											if v549 != 0 {
												return int32(0)
											} else {
												v550 = F_text_to_cstring(m, l0)
												mBase = m.M
												v551 = m.ExcPending
												if v551 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12))) = v550
													F_errmsg(m, int32(_a_F_parse_datetime_13), v12)
													mBase = m.M
													v555 = m.ExcPending
													if v555 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_19), int32(_a_F_parse_datetime_6))
														mBase = m.M
														v560 = m.ExcPending
														if v560 != 0 {
															return int32(0)
														} else {
															v570 = v542
															m.G0 = v12 + int32(96)
															return v570
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1082)
									v570 = v539 - int32(_a_F_parse_datetime_0)
									m.G0 = v12 + int32(96)
									return v570
								}
							} else {
								if v329 != int32(_a_F_parse_datetime_2) {
									v346 = int32(0)
									v347 = F_errsave_start(m, l5)
									mBase = m.M
									v348 = m.ExcPending
									if v348 != 0 {
										return int32(0)
									} else {
										if v347 == int32(0) {
											v570 = v346
											m.G0 = v12 + int32(96)
											return v570
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v353 = m.ExcPending
											if v353 != 0 {
												return int32(0)
											} else {
												v354 = F_text_to_cstring(m, l0)
												mBase = m.M
												v355 = m.ExcPending
												if v355 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v354
													F_errmsg(m, int32(_a_F_parse_datetime_13), v12+int32(16))
													mBase = m.M
													v361 = m.ExcPending
													if v361 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
														mBase = m.M
														v366 = m.ExcPending
														if v366 != 0 {
															return int32(0)
														} else {
															v570 = v346
															m.G0 = v12 + int32(96)
															return v570
														}
													}
												}
											}
										}
									}
								} else {
									v342 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
									if v342 < int32(6) {
										v508 = v342
										v509 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
										v514 = base.B2i32(int32(2) < v508)
										if int32(2) < v508 {
											v515 = int32(_a_F_parse_datetime_15)
										} else {
											v515 = int32(_a_F_parse_datetime_16)
										}
										v516 = v515 + v329
										v521 = base.I32_div_s(v516, int32(4))
										v524 = base.I32_div_s(v516, int32(-100))
										v527 = base.I32_div_s(v516, int32(400))
										if int32(2) < v508 {
											v531 = int32(1)
										} else {
											v531 = int32(13)
										}
										v536 = base.I32_div_s((v531+v508)*int32(_a_F_parse_datetime_17), int32(256))
										v539 = v509 + v516*int32(365) + v521 + v524 + v527 + v536 - int32(_a_F_parse_datetime_18)
										if base.Ui32(int32(2147483494)) <= base.Ui32(v539) {
											v542 = int32(0)
											v543 = F_errsave_start(m, l5)
											mBase = m.M
											v544 = m.ExcPending
											if v544 != 0 {
												return int32(0)
											} else {
												if v543 == int32(0) {
													v570 = v542
													m.G0 = v12 + int32(96)
													return v570
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v549 = m.ExcPending
													if v549 != 0 {
														return int32(0)
													} else {
														v550 = F_text_to_cstring(m, l0)
														mBase = m.M
														v551 = m.ExcPending
														if v551 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12))) = v550
															F_errmsg(m, int32(_a_F_parse_datetime_13), v12)
															mBase = m.M
															v555 = m.ExcPending
															if v555 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_19), int32(_a_F_parse_datetime_6))
																mBase = m.M
																v560 = m.ExcPending
																if v560 != 0 {
																	return int32(0)
																} else {
																	v570 = v542
																	m.G0 = v12 + int32(96)
																	return v570
																}
															}
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1082)
											v570 = v539 - int32(_a_F_parse_datetime_0)
											m.G0 = v12 + int32(96)
											return v570
										}
									} else {
										v346 = int32(0)
										v347 = F_errsave_start(m, l5)
										mBase = m.M
										v348 = m.ExcPending
										if v348 != 0 {
											return int32(0)
										} else {
											if v347 == int32(0) {
												v570 = v346
												m.G0 = v12 + int32(96)
												return v570
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v353 = m.ExcPending
												if v353 != 0 {
													return int32(0)
												} else {
													v354 = F_text_to_cstring(m, l0)
													mBase = m.M
													v355 = m.ExcPending
													if v355 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v354
														F_errmsg(m, int32(_a_F_parse_datetime_13), v12+int32(16))
														mBase = m.M
														v361 = m.ExcPending
														if v361 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
															mBase = m.M
															v366 = m.ExcPending
															if v366 != 0 {
																return int32(0)
															} else {
																v570 = v346
																m.G0 = v12 + int32(96)
																return v570
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
			} else {
				if v38 != 0 {
					if v36&int32(4) != 0 {
						v370 = F_palloc(m, int32(16))
						mBase = m.M
						v371 = m.ExcPending
						if v371 != 0 {
							return int32(0)
						} else {
							v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
							if v372 == int32(1) {
								v375 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v375
								v377 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
								v379 = v12 + int32(52)
								v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
								v381 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
								v382 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v370)+8)) = v375
								v385 = int32(60)
								*(*int64)(unsafe.Add(mBase, uint32(v370))) = base.I64_extend_i32_s(v377) + base.I64_extend_i32_s(v380+(v381+v382*v385)*v385)*int64(1000000)
								v413 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								if base.Ui32(v413) <= base.Ui32(int32(6)) {
									v420 = v413 << (uint(int32(3)) % 32)
									v421 = *(*int64)(unsafe.Add(mBase, uint32(v420)+uint32(_c_F_parse_datetime[0])))
									v422 = *(*int64)(unsafe.Add(mBase, uint32(v420)+uint32(_c_F_parse_datetime[1])))
									v423 = *(*int64)(unsafe.Add(mBase, uint32(v370)))
									if int64(0) <= v423 {
										v426 = v422 + v423
										v427 = base.I64_rem_s(v426, v421)
										v433 = v426 - v427
									} else {
										v429 = v422 - v423
										v430 = base.I64_rem_s(v429, v421)
										v433 = v430 - v429
									}
									*(*int64)(unsafe.Add(mBase, uint32(v370))) = v433
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1266)
								v570 = v370
								m.G0 = v12 + int32(96)
								return v570
							} else {
								v396 = int32(0)
								v397 = F_errsave_start(m, l5)
								mBase = m.M
								v398 = m.ExcPending
								if v398 != 0 {
									return int32(0)
								} else {
									if v397 == int32(0) {
										v570 = v396
										m.G0 = v12 + int32(96)
										return v570
									} else {
										F_errcode(m, int32(117440642))
										mBase = m.M
										v403 = m.ExcPending
										if v403 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_parse_datetime_20), int32(0))
											mBase = m.M
											v407 = m.ExcPending
											if v407 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_21), int32(_a_F_parse_datetime_6))
												mBase = m.M
												v412 = m.ExcPending
												if v412 != 0 {
													return int32(0)
												} else {
													v570 = v396
													m.G0 = v12 + int32(96)
													return v570
												}
											}
										}
									}
								}
							}
						}
					} else {
						v442 = v12 + int32(24)
						v443 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+40)))
						v445 = v12 + int32(52)
						v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
						v447 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
						v448 = *(*int32)(unsafe.Add(mBase, uint32(v445)+8))
						v449 = int32(60)
						*(*int64)(unsafe.Add(mBase, uint32(v442))) = v443 + base.I64_extend_i32_s(v446+(v447+v448*v449)*v449)*int64(1000000)
						v460 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						if base.Ui32(v460) <= base.Ui32(int32(6)) {
							v467 = v460 << (uint(int32(3)) % 32)
							v468 = *(*int64)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_parse_datetime[0])))
							v469 = *(*int64)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_parse_datetime[1])))
							v470 = *(*int64)(unsafe.Add(mBase, uint32(v442)))
							if int64(0) <= v470 {
								v473 = v469 + v470
								v474 = base.I64_rem_s(v473, v468)
								v480 = v473 - v474
							} else {
								v476 = v469 - v470
								v477 = base.I64_rem_s(v476, v468)
								v480 = v477 - v476
							}
							*(*int64)(unsafe.Add(mBase, uint32(v442))) = v480
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1083)
						v488 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
						v489 = F_Int64GetDatum(m, v488)
						mBase = m.M
						v490 = m.ExcPending
						if v490 != 0 {
							return int32(0)
						} else {
							v570 = v489
							m.G0 = v12 + int32(96)
							return v570
						}
					}
				} else {
					v491 = int32(0)
					v492 = F_errsave_start(m, l5)
					mBase = m.M
					v493 = m.ExcPending
					if v493 != 0 {
						return int32(0)
					} else {
						if v492 == int32(0) {
							v570 = v491
							m.G0 = v12 + int32(96)
							return v570
						} else {
							F_errcode(m, int32(117440642))
							mBase = m.M
							v498 = m.ExcPending
							if v498 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_parse_datetime_22), int32(0))
								mBase = m.M
								v502 = m.ExcPending
								if v502 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_23), int32(_a_F_parse_datetime_6))
									mBase = m.M
									v507 = m.ExcPending
									if v507 != 0 {
										return int32(0)
									} else {
										v570 = v491
										m.G0 = v12 + int32(96)
										return v570
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
func F_parse_ident_line(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != 0 {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
		v21 = v19
	} else {
		v21 = int32(0)
	}
	v22 = int32(16)
	v23 = l0 + v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = F_palloc0(m, v22)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = v17
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
		if int32(2) <= v32 {
			v35 = int32(0)
			v37 = F_errstart(m, l1, v35)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if v37 != 0 {
					F_errcode(m, int32(22))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_parse_ident_line_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v24
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
								F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2769), int32(_a_F_parse_ident_line_3))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										v60 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v23))) = v60
											v262 = v35
											m.G0 = v15 + int32(80)
											return v262
										}
									}
								}
							}
						}
					}
				} else {
					v60 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v23))) = v60
						v262 = v35
						m.G0 = v15 + int32(80)
						return v262
					}
				}
			}
		} else {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v66 = F_pstrdup(m, v65)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v66
				v70 = v21 + int32(4)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
				if base.Ui32(v72+v73<<(uint(int32(2))%32)) <= base.Ui32(v70) {
					v78 = int32(0)
					v80 = F_errstart(m, l1, v78)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						if v80 != 0 {
							F_errcode(m, int32(22))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_parse_ident_line_4), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_set_errcontext_domain(m, int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v24
										*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v17
										F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15-int32(-64))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2775), int32(_a_F_parse_ident_line_3))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v105 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v23))) = v105
													v262 = v78
													m.G0 = v15 + int32(80)
													return v262
												}
											}
										}
									}
								}
							}
						} else {
							v105 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v23))) = v105
								v262 = v78
								m.G0 = v15 + int32(80)
								return v262
							}
						}
					}
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
					if int32(2) <= v109 {
						v112 = int32(0)
						v114 = F_errstart(m, l1, v112)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							if v114 != 0 {
								F_errcode(m, int32(22))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_parse_ident_line_0), int32(0))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										F_set_errcontext_domain(m, int32(0))
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v24
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v17
											F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15+int32(16))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2777), int32(_a_F_parse_ident_line_3))
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													v139 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v23))) = v139
														v262 = v112
														m.G0 = v15 + int32(80)
														return v262
													}
												}
											}
										}
									}
								}
							} else {
								v139 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23))) = v139
									v262 = v112
									m.G0 = v15 + int32(80)
									return v262
								}
							}
						}
					} else {
						v142 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
						v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+4)))
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
						v146 = F_strlen(m, v145)
						mBase = m.M
						v149 = F_palloc0(m, v146+int32(13))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v149)+8)) = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v149)+4)) = uint8(v144)
							v155 = v149 + int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(v149))) = v155
							v158 = v146 + int32(1)
							if v158 != 0 {
								base.MemoryCopy(m, v155, v145, v158)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v149
							v162 = v21 + int32(8)
							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
							v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
							if base.Ui32(v164+v165<<(uint(int32(2))%32)) <= base.Ui32(v162) {
								v170 = int32(0)
								v172 = F_errstart(m, l1, v170)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return int32(0)
								} else {
									if v172 != 0 {
										F_errcode(m, int32(22))
										mBase = m.M
										v176 = m.ExcPending
										if v176 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_parse_ident_line_4), int32(0))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return int32(0)
											} else {
												F_set_errcontext_domain(m, int32(0))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v24
													*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v17
													F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15+int32(48))
													mBase = m.M
													v190 = m.ExcPending
													if v190 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2785), int32(_a_F_parse_ident_line_3))
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int32(0)
														} else {
															v197 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
															mBase = m.M
															v198 = m.ExcPending
															if v198 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v23))) = v197
																v262 = v170
																m.G0 = v15 + int32(80)
																return v262
															}
														}
													}
												}
											}
										}
									} else {
										v197 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v23))) = v197
											v262 = v170
											m.G0 = v15 + int32(80)
											return v262
										}
									}
								}
							} else {
								v200 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
								v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
								if int32(2) <= v201 {
									v204 = int32(0)
									v206 = F_errstart(m, l1, v204)
									mBase = m.M
									v207 = m.ExcPending
									if v207 != 0 {
										return int32(0)
									} else {
										if v206 != 0 {
											F_errcode(m, int32(22))
											mBase = m.M
											v210 = m.ExcPending
											if v210 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_parse_ident_line_0), int32(0))
												mBase = m.M
												v214 = m.ExcPending
												if v214 != 0 {
													return int32(0)
												} else {
													F_set_errcontext_domain(m, int32(0))
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v24
														*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v17
														F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15+int32(32))
														mBase = m.M
														v224 = m.ExcPending
														if v224 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2787), int32(_a_F_parse_ident_line_3))
															mBase = m.M
															v229 = m.ExcPending
															if v229 != 0 {
																return int32(0)
															} else {
																v231 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
																mBase = m.M
																v232 = m.ExcPending
																if v232 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v231
																	v262 = v204
																	m.G0 = v15 + int32(80)
																	return v262
																}
															}
														}
													}
												}
											}
										} else {
											v231 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
											mBase = m.M
											v232 = m.ExcPending
											if v232 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v23))) = v231
												v262 = v204
												m.G0 = v15 + int32(80)
												return v262
											}
										}
									}
								} else {
									v234 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
									v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
									v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+4)))
									v238 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
									v239 = F_strlen(m, v238)
									mBase = m.M
									v242 = F_palloc0(m, v239+int32(13))
									mBase = m.M
									v243 = m.ExcPending
									if v243 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v242)+4)) = uint8(v236)
										v248 = v242 + int32(12)
										*(*int32)(unsafe.Add(mBase, uint32(v242))) = v248
										v251 = v239 + int32(1)
										if v251 != 0 {
											base.MemoryCopy(m, v248, v238, v251)
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v242
										v254 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
										v255 = F_regcomp_auth_token(m, v254, v24, v17, v23, l1)
										mBase = m.M
										v256 = m.ExcPending
										if v256 != 0 {
											return int32(0)
										} else {
											if v255 != 0 {
												v262 = int32(0)
												m.G0 = v15 + int32(80)
												return v262
											} else {
												v258 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
												v259 = F_regcomp_auth_token(m, v258, v24, v17, v23, l1)
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
													return int32(0)
												} else {
													if v259 != 0 {
														v261 = int32(0)
													} else {
														v261 = v26
													}
													v262 = v261
													m.G0 = v15 + int32(80)
													return v262
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
}
func F_pathkeys_count_contained_in(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	if l0 == l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v76
	return int32(1)
L2:
	;
	if l0 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	return int32(1)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	return int32(1)
L7:
	;
	goto L8
L8:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23
	return v23
L10:
	;
	goto L11
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = int32(0)
	if v28 < v27 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = v27
	goto L14
L13:
	;
	v31 = v28
	goto L14
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = int32(0)
	goto L15
L15:
	;
	if v37 < v32 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = v44 + v37<<(uint(int32(2))%32)
	goto L19
L18:
	;
	v48 = int32(0)
	goto L19
L19:
	;
	if v37 == v31 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31
	return base.B2i32(v48 == int32(0))
L21:
	;
	goto L22
L22:
	;
	v55 = base.B2i32(v48 == int32(0))
	if v48 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
	return v55
L24:
	;
	goto L25
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v60 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
	return v55
L27:
	;
	goto L28
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60+v37<<(uint(int32(2))%32))))
	if v65 != v69 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
	return int32(0)
L30:
	;
	v37 = v37 + int32(1)
	goto L15
}
func F_pgarch_waken_stop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	*(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[1]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[2]))
	if v17 == v13 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[3]))
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v47 = F_pgmem_kill(m, v13, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v21 + int32(16)
	goto L1
L10:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v27)
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[4]))
	v35 = F_write(m, v31, v21+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[5]))
	if v39 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_pglz_decompress_datum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_palloc(m, v3&int32(1073741823)+int32(4))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = int32(8)
	v13 = l0 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = int32(base.Ui32(v14)>>(uint(int32(2))%32)) - v12
	v20 = v8 + int32(4)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = v21 & int32(1073741823)
	v25 = int32(0)
	v32 = v20 + v23
	v33 = v13 + v18
	if base.B2i32(v18 <= v25)|base.B2i32(v23 <= v25) == v25 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v208 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L4:
	;
	goto L3
L5:
	;
	goto L45
L6:
	;
	v41 = v13
	v44 = v20
	goto L9
L7:
	;
	goto L8
L8:
	;
	v183 = v13
	v186 = v20
	goto L5
L9:
	;
	v54 = v41 + int32(1)
	if base.Ui32(v33) <= base.Ui32(v54) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v183 = v169
	v186 = v172
	goto L5
L11:
	;
	if base.Ui32(v33) <= base.Ui32(v169) {
		v183 = v169
		v186 = v172
		goto L5
	} else {
		goto L43
	}
L12:
	;
	v169 = v54
	v172 = v44
	goto L11
L13:
	;
	goto L14
L14:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v59 = v54
	v61 = v44
	v67 = v56
	v68 = int32(0)
	goto L15
L15:
	;
	if v67&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v169 = v145
	v172 = v157
	goto L11
L17:
	;
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v68))|base.B2i32(base.Ui32(v33) <= base.Ui32(v145)) != 0 {
		v169 = v145
		v172 = v157
		goto L11
	} else {
		goto L41
	}
L18:
	;
	v72 = int32(-1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v77 = v73&int32(15) + int32(3)
	if v77 != int32(18) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v139)
	v141 = int32(1)
	v145 = v59 + v141
	v157 = v61 + v141
	goto L17
L21:
	;
	v87 = v77
	v88 = v59 + int32(2)
	goto L23
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	v87 = v82 + int32(18)
	v88 = v59 + int32(3)
	goto L23
L23:
	;
	if base.Ui32(v33) < base.Ui32(v88) {
		v208 = v72
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	v95 = v90 | v73<<(uint(int32(4))%32)&int32(3840)
	if base.B2i32(v95 == int32(0))|base.B2i32(v61-v20 < v95) != 0 {
		v208 = v72
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v101 = v32 - v61
	if v87 < v101 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v103 = v87
	goto L28
L27:
	;
	v103 = v101
	goto L28
L28:
	;
	if v95 < v103 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v106 = v95
	v108 = v61
	v110 = v103
	goto L32
L30:
	;
	v125 = v95
	v127 = v61
	v129 = v103
	goto L31
L31:
	;
	if v129 != 0 {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	if v106 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v125 = v122
	v127 = v119
	v129 = v120
	goto L31
L34:
	;
	base.MemoryCopy(m, v108, v108-v106, v106)
	goto L36
L35:
	;
	goto L36
L36:
	;
	v119 = v106 + v108
	v120 = v110 - v106
	v122 = v106 << (uint(int32(1)) % 32)
	if v122 < v120 {
		v106 = v122
		v108 = v119
		v110 = v120
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	base.MemoryCopy(m, v127, v127-v125, v129)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v145 = v88
	v157 = v127 + v129
	goto L17
L41:
	;
	v162 = int32(1)
	if base.Ui32(v157) < base.Ui32(v32) {
		v59 = v145
		v61 = v157
		v67 = int32(base.Ui32(v67&int32(254)) >> (uint(v162) % 32))
		v68 = v68 + v162
		goto L15
	} else {
		goto L42
	}
L42:
	;
	goto L16
L43:
	;
	if base.Ui32(v172) < base.Ui32(v32) {
		v41 = v169
		v44 = v172
		goto L9
	} else {
		goto L44
	}
L44:
	;
	goto L10
L45:
	;
	if base.B2i32(v183 != v33)|base.B2i32(v186 != v32) != 0 {
		v208 = int32(-1)
		goto L4
	} else {
		goto L48
	}
L47:
	;
	v208 = v186 - v20
	goto L4
L48:
	;
	goto L47
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v208<<(uint(int32(2))%32) + int32(16)
	return v8
L52:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_pglz_decompress_datum_0), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_pglz_decompress_datum_1), int32(98), int32(_a_F_pglz_decompress_datum_2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstatginindex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pgstatginindex_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstatginindex_1), int32(489), int32(_a_F_pgstatginindex_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
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
			v26 = F_pgstatginindex_internal(m, v3, l0)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v26
			}
		}
	}
}
func F_pgstathashindex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int64
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int64
	_ = v238
	var v246 int64
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v378 int64
	_ = v378
	var v379 int64
	_ = v379
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int64
	_ = v403
	var v411 int64
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v439 int64
	_ = v439
	var v440 int64
	_ = v440
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v482 int64
	_ = v482
	var v483 int64
	_ = v483
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int64
	_ = v567
	var v568 int64
	_ = v568
	var v569 int64
	_ = v569
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v599 int32
	_ = v599
	var v604 int64
	_ = v604
	var v605 int64
	_ = v605
	var v606 int64
	_ = v606
	var v609 int64
	_ = v609
	var v610 int64
	_ = v610
	var v613 int64
	_ = v613
	var v616 int32
	_ = v616
	var v621 int64
	_ = v621
	var v632 float64
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	v2 = int32(0)
	v21 = int64(0)
	v30 = m.G0
	v32 = v30 - int32(112)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+56)) = v21
	v38 = F_relation_open(m, v34, int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L6
	} else {
		goto L122
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L6
	} else {
		goto L114
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L110
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L6
	} else {
		goto L106
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L6
	} else {
		goto L102
	}
L6:
	;
	return int32(0)
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	if v43 != int32(105) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+84))
	if v46 != int32(405) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)))
	if v49 == int32(116) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+24)))
	if v52 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+192))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+18)))
	if v56 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v62 = F__hash_getbuf(m, v38, int32(0), int32(1), int32(8))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L16
	}
L15:
	;
	v82 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v81)+42)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+28))
	F_UnlockReleaseBuffer(m, v62)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L20
	}
L16:
	;
	if v62 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[0]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+(v62^int32(-1))<<(uint(int32(2))%32))))
	v81 = v73
	goto L15
L18:
	;
	goto L19
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[1]))
	v81 = v75 + v62<<(uint(int32(13))%32) + int32(-8192)
	goto L15
L20:
	;
	v87 = F_RelationGetNumberOfBlocksInFork(m, v38, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v89 = int32(1)
	v91 = F_GetAccessStrategy(m, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32(v87) < base.Ui32(int32(2)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v599 = v2
	v604 = v21
	v605 = v21
	v606 = v21
	v609 = v21
	v610 = v21
	v613 = int64(0)
	goto L25
L24:
	;
	v106 = v2
	v108 = v89
	v111 = v2
	v112 = v2
	v113 = v2
	v116 = v21
	v117 = v21
	v118 = v21
	goto L26
L25:
	;
	F_relation_close(m, v38, int32(1))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L6
	} else {
		goto L86
	}
L26:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[2]))
	if v126 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v599 = v562
	v604 = v567
	v605 = v568
	v606 = v569
	v609 = base.I64_extend_i32_u(v564)
	v610 = base.I64_extend_i32_u(v563)
	v613 = base.I64_extend_i32_u(v557)
	goto L25
L28:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v129 = int32(0)
	v131 = F_ReadBufferExtended(m, v38, v129, v108, v129, v91)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	F_LockBuffer(m, v131, int32(1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	if v131 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	F_UnlockReleaseBuffer(m, v131)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L6
	} else {
		goto L84
	}
L35:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+14)))
	if v154 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[0]))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139+(v131^int32(-1))<<(uint(int32(2))%32))))
	v153 = v145
	goto L35
L37:
	;
	goto L38
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[1]))
	v153 = v147 + v131<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v557 = v106 + int32(1)
	v562 = v111
	v563 = v112
	v564 = v113
	v567 = v116
	v568 = v117
	v569 = v118
	goto L34
L40:
	;
	goto L41
L41:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+19)))
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+16)))
	if (v159<<(uint(int32(8))%32)-v162)&int32(_a_F_pgstathashindex_0) != int32(16) {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v168 = v162 + v153
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+12)))
	switch v169 & int32(15) {
	case 0:
		goto L45
	case 1:
		goto L46
	case 2:
		goto L47
	default:
		goto L44
	case 4:
		goto L43
	}
L43:
	;
	v557 = v106
	v562 = v111 + int32(1)
	v563 = v112
	v564 = v113
	v567 = v116
	v568 = v117
	v569 = v118
	goto L34
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L6
	} else {
		goto L76
	}
L45:
	;
	v557 = v106 + int32(1)
	v562 = v111
	v563 = v112
	v564 = v113
	v567 = v116
	v568 = v117
	v569 = v118
	goto L34
L46:
	;
	v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+12)))
	if base.Ui32(v337) < base.Ui32(int32(25)) {
		v482 = v116
		v483 = v117
		goto L62
	} else {
		goto L63
	}
L47:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+12)))
	if base.Ui32(v172) < base.Ui32(int32(25)) {
		v317 = v116
		v318 = v117
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+14)))
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+12)))
	v330 = v328 - v329
	v331 = int32(0)
	if v331 < v330 {
		goto L59
	} else {
		goto L60
	}
L49:
	;
	v178 = int32(base.Ui32(v172+int32(_a_F_pgstathashindex_1)) >> (uint(int32(2)) % 32))
	v180 = v178 & int32(_a_F_pgstathashindex_0)
	if v180 == int32(0) {
		v317 = v116
		v318 = v117
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v183 = int32(1)
	v185 = v153 + int32(20)
	if v180 != v183 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v195 = v183
	v198 = int32(0)
	v213 = v116
	v214 = v117
	goto L54
L52:
	;
	v256 = v183
	v274 = v116
	v275 = v117
	goto L53
L53:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v185+v256<<(uint(int32(2))%32))))
	v287 = int32(_a_F_pgstathashindex_2)
	v288 = v286 & v287
	v317 = v274 + base.I64_extend_i32_u(base.B2i32(v288 == v287))
	v318 = v275 + base.I64_extend_i32_u(base.B2i32(v288 != v287))
	goto L48
L54:
	;
	v222 = int32(2)
	v224 = v185 + v195<<(uint(v222)%32)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v226 = int32(_a_F_pgstathashindex_2)
	v227 = v225 & v226
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v234 = v232 & v226
	v238 = v213 + base.I64_extend_i32_u(base.B2i32(v227 == v226)) + base.I64_extend_i32_u(base.B2i32(v234 == v226))
	v246 = base.I64_extend_i32_u(base.B2i32(v234 != v226)) + (v214 + base.I64_extend_i32_u(base.B2i32(v227 != v226)))
	v248 = v195 + v222
	v250 = v198 + v222
	if v250 != v178&int32(_a_F_pgstathashindex_3) {
		v195 = v248
		v198 = v250
		v213 = v238
		v214 = v246
		goto L54
	} else {
		goto L56
	}
L55:
	;
	if v178&int32(1) == int32(0) {
		v317 = v238
		v318 = v246
		goto L48
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v256 = v248
	v274 = v238
	v275 = v246
	goto L53
L58:
	;
	v557 = v106
	v562 = v111
	v563 = v112 + int32(1)
	v564 = v113
	v567 = v317
	v568 = v318
	v569 = v118 + base.I64_extend_i32_u(v334)
	goto L34
L59:
	;
	v334 = v330
	goto L61
L60:
	;
	v334 = v331
	goto L61
L61:
	;
	goto L58
L62:
	;
	v493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+14)))
	v494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+12)))
	v495 = v493 - v494
	v496 = int32(0)
	if v496 < v495 {
		goto L73
	} else {
		goto L74
	}
L63:
	;
	v343 = int32(base.Ui32(v337+int32(_a_F_pgstathashindex_1)) >> (uint(int32(2)) % 32))
	v345 = v343 & int32(_a_F_pgstathashindex_0)
	if v345 == int32(0) {
		v482 = v116
		v483 = v117
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v348 = int32(1)
	v350 = v153 + int32(20)
	if v345 != v348 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v360 = v348
	v363 = int32(0)
	v378 = v116
	v379 = v117
	goto L68
L66:
	;
	v421 = v348
	v439 = v116
	v440 = v117
	goto L67
L67:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v350+v421<<(uint(int32(2))%32))))
	v452 = int32(_a_F_pgstathashindex_2)
	v453 = v451 & v452
	v482 = v439 + base.I64_extend_i32_u(base.B2i32(v453 == v452))
	v483 = v440 + base.I64_extend_i32_u(base.B2i32(v453 != v452))
	goto L62
L68:
	;
	v387 = int32(2)
	v389 = v350 + v360<<(uint(v387)%32)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v391 = int32(_a_F_pgstathashindex_2)
	v392 = v390 & v391
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v399 = v397 & v391
	v403 = v378 + base.I64_extend_i32_u(base.B2i32(v392 == v391)) + base.I64_extend_i32_u(base.B2i32(v399 == v391))
	v411 = base.I64_extend_i32_u(base.B2i32(v399 != v391)) + (v379 + base.I64_extend_i32_u(base.B2i32(v392 != v391)))
	v413 = v360 + v387
	v415 = v363 + v387
	if v415 != v343&int32(_a_F_pgstathashindex_3) {
		v360 = v413
		v363 = v415
		v378 = v403
		v379 = v411
		goto L68
	} else {
		goto L70
	}
L69:
	;
	if v343&int32(1) == int32(0) {
		v482 = v403
		v483 = v411
		goto L62
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	v421 = v413
	v439 = v403
	v440 = v411
	goto L67
L72:
	;
	v557 = v106
	v562 = v111
	v563 = v112
	v564 = v113 + int32(1)
	v567 = v482
	v568 = v483
	v569 = v118 + base.I64_extend_i32_u(v499)
	goto L34
L73:
	;
	v499 = v495
	goto L75
L74:
	;
	v499 = v496
	goto L75
L75:
	;
	goto L72
L76:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+12)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	if v131 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v512 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v511
	F_errmsg(m, int32(_a_F_pgstathashindex_4), v32)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L6
	} else {
		goto L82
	}
L79:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[3]))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v516+(v131^int32(-1))<<(uint(int32(6))%32))+16))
	v531 = v522
	goto L78
L80:
	;
	goto L81
L81:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[4]))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v524+v131<<(uint(int32(6))%32)+int32(-64))+16))
	v531 = v530
	goto L78
L82:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(688), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	v579 = v108 + int32(1)
	if v579 != v87 {
		v106 = v557
		v108 = v579
		v111 = v562
		v112 = v563
		v113 = v564
		v116 = v567
		v117 = v568
		v118 = v569
		goto L26
	} else {
		goto L85
	}
L85:
	;
	goto L27
L86:
	;
	v621 = base.I64_extend_i32_u(v87+(v599^int32(-1))) * v82
	if v621 == int64(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v632 = float64(0)
	goto L89
L88:
	;
	v632 = base.F64_div(base.F64_mul(base.F64_convert_i64_u(v606+v82*v613), float64(100)), base.F64_convert_i64_u(v621))
	goto L89
L89:
	;
	v636 = F_get_call_result_type(m, l0, int32(0), v32+int32(108))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L6
	} else {
		goto L90
	}
L90:
	;
	if v636 != int32(1) {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v32)+108))
	v641 = F_BlessTupleDesc(m, v640)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v641
	v645 = F_Int64GetDatum(m, v610)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v645
	v648 = F_Int64GetDatum(m, v609)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v648
	v652 = F_Int64GetDatum(m, base.I64_extend_i32_u(v599))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v652
	v655 = F_Int64GetDatum(m, v613)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = v655
	v658 = F_Int64GetDatum(m, v605)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v658
	v661 = F_Int64GetDatum(m, v604)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v661
	v664 = F_Float8GetDatum(m, v632)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v664
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v32)+108))
	v672 = F_heap_form_tuple(m, v667, v32-int32(-64), v32+int32(56))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v672)+16))
	v675 = F_HeapTupleHeaderGetDatum(m, v674)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	m.G0 = v32 + int32(112)
	return v675
L102:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v688 + int32(4)
	F_errmsg(m, int32(_a_F_pgstathashindex_7), v32+int32(48))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(606), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L6
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(_a_F_pgstathashindex_8), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(616), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L6
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
	F_errcode(m, int32(325))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v725 + int32(4)
	F_errmsg(m, int32(_a_F_pgstathashindex_9), v32+int32(32))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(623), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L6
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	if v131 < int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v765
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v746 + int32(4)
	F_errmsg(m, int32(_a_F_pgstathashindex_10), v32+int32(16))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L6
	} else {
		goto L120
	}
L117:
	;
	v750 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[3]))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v750+(v131^int32(-1))<<(uint(int32(6))%32))+16))
	v765 = v756
	goto L116
L118:
	;
	goto L119
L119:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[4]))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v758+v131<<(uint(int32(6))%32)+int32(-64))+16))
	v765 = v764
	goto L116
L120:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(660), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errmsg_internal(m, int32(_a_F_pgstathashindex_11), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(715), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstatindex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_superuser(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			if v8 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pgstatindex_0), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pgstatindex_1), int32(151), int32(_a_F_pgstatindex_2))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
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
				v28 = F_textToQualifiedNameList(m, v4)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = F_makeRangeVarFromNameList(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v33 = F_relation_openrv(m, v30, int32(1))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = F_pgstatindex_impl(m, v33, l0)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								return v35
							}
						}
					}
				}
			}
		}
	}
}
func F_pgstattuple_v1_5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_textToQualifiedNameList(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = F_makeRangeVarFromNameList(m, v7)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v12 = F_relation_openrv(m, v9, int32(1))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					v14 = F_pgstat_relation(m, v12, l0)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int32(0)
					} else {
						return v14
					}
				}
			}
		}
	}
}
func F_pkt_stream_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(70368744177664)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5
		return int32(_a_F_pkt_stream_init_0)
	}
}
func F_pktreader_free(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	base.MemoryFill(m, l0, int32(0), int32(8))
	F_pfree(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_placeChar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	if l0 == int32(0) {
		v10 = F_palloc0(m, int32(3072))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = v10
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v18 = v14 + v15*int32(12)
			if l2 <= int32(1) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				if v21 != 0 {
					v24 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if v24 == int32(0) {
							return v14
						} else {
							F_errcode(m, int32(22))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_placeChar_0), int32(0))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_placeChar_1), int32(74), int32(_a_F_placeChar_2))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										return v14
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l4
					v42 = F_palloc(m, l4)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v42
						if l4 == int32(0) {
						} else {
							base.MemoryCopy(m, v42, l3, l4)
						}
						return v14
					}
				}
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v51 = int32(1)
				v55 = F_placeChar(m, v50, l1+v51, l2-v51, l3, l4)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v55
					return v14
				}
			}
		}
	} else {
		v14 = l0
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v18 = v14 + v15*int32(12)
		if l2 <= int32(1) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			if v21 != 0 {
				v24 = F_errstart(m, int32(19), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						return v14
					} else {
						F_errcode(m, int32(22))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_placeChar_0), int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_placeChar_1), int32(74), int32(_a_F_placeChar_2))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									return v14
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l4
				v42 = F_palloc(m, l4)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v42
					if l4 == int32(0) {
					} else {
						base.MemoryCopy(m, v42, l3, l4)
					}
					return v14
				}
			}
		} else {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v51 = int32(1)
			v55 = F_placeChar(m, v50, l1+v51, l2-v51, l3, l4)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v55
				return v14
			}
		}
	}
}
func F_points_box(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v20 float64
	_ = v20
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v31 float64
	_ = v31
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v43 float64
	_ = v43
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v55 float64
	_ = v55
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_palloc(m, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		if base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			if base.F64_lt(v14, v20) != 0 {
				v28 = v20
				v29 = v8
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v20)&int64(9223372036854775807)) {
					v28 = v20
					v29 = v8
				} else {
					v28 = v14
					v29 = v7
				}
			}
		} else {
			v28 = v14
			v29 = v7
		}
		*(*float64)(unsafe.Add(mBase, uint32(v10))) = v28
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
		*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v31
		v34 = v7 + int32(8)
		v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		if base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v43 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
			if base.F64_gt(v43, v35)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v43)&int64(9223372036854775807))) != 0 {
				v52 = v43
				v53 = v8 + int32(8)
			} else {
				v52 = v35
				v53 = v34
			}
		} else {
			v52 = v35
			v53 = v34
		}
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v52
		v55 = *(*float64)(unsafe.Add(mBase, uint32(v53)))
		*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v55
		return v10
	}
}
func F_pqinitmask(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	*(*int64)(unsafe.Add(mBase, _c_F_pqinitmask[0])) = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = int64(-15032385537)
	*(*int64)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = int64(-15032385537)
	v31 = int32(_a_F_pqinitmask_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v32 & base.I32_rotl(int32(-2), int32(4))
	v59 = int32(_a_F_pqinitmask_1)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v60 & base.I32_rotl(int32(-2), int32(4))
	v87 = int32(_a_F_pqinitmask_0)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v88 & base.I32_rotl(int32(-2), int32(5))
	v115 = int32(_a_F_pqinitmask_1)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v116 & base.I32_rotl(int32(-2), int32(5))
	v143 = int32(_a_F_pqinitmask_0)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v144 & base.I32_rotl(int32(-2), int32(3))
	v171 = int32(_a_F_pqinitmask_1)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v172 & base.I32_rotl(int32(-2), int32(3))
	v199 = int32(_a_F_pqinitmask_0)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v200 & base.I32_rotl(int32(-2), int32(7))
	v227 = int32(_a_F_pqinitmask_1)
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v228 & base.I32_rotl(int32(-2), int32(7))
	v255 = int32(_a_F_pqinitmask_0)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v256 & base.I32_rotl(int32(-2), int32(10))
	v283 = int32(_a_F_pqinitmask_1)
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v284 & base.I32_rotl(int32(-2), int32(10))
	v311 = int32(_a_F_pqinitmask_0)
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v312 & base.I32_rotl(int32(-2), int32(6))
	v339 = int32(_a_F_pqinitmask_1)
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v340 & base.I32_rotl(int32(-2), int32(6))
	v367 = int32(_a_F_pqinitmask_0)
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v368 & base.I32_rotl(int32(-2), int32(30))
	v395 = int32(_a_F_pqinitmask_1)
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v396 & base.I32_rotl(int32(-2), int32(30))
	v423 = int32(_a_F_pqinitmask_0)
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v424 & base.I32_rotl(int32(-2), int32(17))
	v451 = int32(_a_F_pqinitmask_1)
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v452 & base.I32_rotl(int32(-2), int32(17))
	v479 = int32(_a_F_pqinitmask_1)
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v480 & base.I32_rotl(int32(-2), int32(2))
	v507 = int32(_a_F_pqinitmask_1)
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v508 & base.I32_rotl(int32(-2), int32(14))
	v535 = int32(_a_F_pqinitmask_1)
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v536 & base.I32_rotl(int32(-2), int32(13))
	return
}
func F_pqsignal_be(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	switch l1 + int32(2) {
	case 0, 2:
		v16 = l1
	default:
		*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_pqsignal_be[0]))) = l1
		v16 = int32(_a_F_pqsignal_be_0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v6+int32(16)))) = int64(0)
	if l0 == int32(17) {
		v26 = int32(268435457)
	} else {
		v26 = int32(268435456)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v26
	v29 = v6 + int32(12)
	if base.Ui32(int32(65)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _c_F_pqsignal_be[1])) = int32(28)
	} else {
		if v29 != 0 {
			v48 = l0 * int32(20)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_pqsignal_be[2]))) = v49
			v51 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_pqsignal_be[3]))) = v51
			v53 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
			*(*int64)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_pqsignal_be[4]))) = v53
		} else {
		}
	}
	m.G0 = v6 + int32(32)
	return
}
func F_pre_format_elog_string(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_pre_format_elog_string[0])) = l0
	return
}
func F_predicatelock_twophase_recover(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v12 {
	case 0:
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
		v18 = F_LWLockAcquire(m, v14+int32(3584), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			if base.B2i32(v22 == int32(0))|base.B2i32(v22 == v21) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v265 = m.ExcPending
				if v265 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_predicatelock_twophase_recover_0))
					mBase = m.M
					v268 = m.ExcPending
					if v268 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_predicatelock_twophase_recover_1), int32(0))
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_predicatelock_twophase_recover_2), int32(_a_F_predicatelock_twophase_recover_3), int32(_a_F_predicatelock_twophase_recover_4))
							mBase = m.M
							v277 = m.ExcPending
							if v277 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
				v33 = v21 + int32(8)
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
				if v34 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v33
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v33
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v40
				*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = v22
				v45 = v22 + int32(-64)
				*(*int32)(unsafe.Add(mBase, uint32(v45))) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = int64(-4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v22-int32(60)))) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v22-int32(48)))) = int64(-1)
				v61 = int64(1)
				*(*int64)(unsafe.Add(mBase, uint32(v22-int32(56)))) = v61
				v64 = v22 + int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v64
				*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v64
				*(*int64)(unsafe.Add(mBase, uint32(v22-int32(40)))) = v61
				*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v22-int32(8)))) = int64(0)
				v79 = v22 - int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v22-int32(12)))) = v79
				*(*int32)(unsafe.Add(mBase, uint32(v79))) = v79
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v82
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v84
				if v84&int32(32) != 0 {
					v93 = v84
				} else {
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v88 + int32(1)
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
					v93 = v92
				}
				*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v93 | int32(1536)
				v100 = v22 - int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v22-int32(20)))) = v100
				v105 = v22 - int32(32)
				*(*int32)(unsafe.Add(mBase, uint32(v22-int32(28)))) = v105
				*(*int32)(unsafe.Add(mBase, uint32(v100))) = v100
				*(*int32)(unsafe.Add(mBase, uint32(v105))) = v105
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l0
				v111 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[2]))
				v117 = F_hash_search(m, v111, v10+int32(12), int32(1), v10+int32(11))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v45
					v121 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
					v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
					if v122 != 0 {
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v123))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v122)) == int32(0) {
							v135 = base.B2i32(base.Ui32(v123) < base.Ui32(v122))
						} else {
							v135 = base.B2i32(int32(0) < v122-v123)
						}
						v137 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
						if v135 == int32(0) {
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
							v207 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
							if v206 != v207 {
							} else {
								v209 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v209 + int32(1)
							}
							v217 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
							F_LWLockRelease(m, v217+int32(3584))
							mBase = m.M
							v221 = m.ExcPending
							if v221 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						} else {
							v140 = v137
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v140)+16)) = v141
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
							v147 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
							v151 = F_LWLockAcquire(m, v147+int32(_a_F_predicatelock_twophase_recover_5), int32(0))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return
							} else {
								if v145 == int32(0) {
									v156 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
									*(*int64)(unsafe.Add(mBase, uint32(v156)+8)) = int64(0)
								} else {
									v161 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])))
									if v161 == int32(1) {
										v166 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[5]))
										v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+316))
										v169 = base.B2i32(v167 != int32(2))
										*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])) = uint8(v169)
										v171 = v169
									} else {
										v171 = int32(0)
									}
									v173 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
									if v171 == int32(0) {
										v196 = v173
										*(*int32)(unsafe.Add(mBase, uint32(v196)+12)) = v145
									} else {
										v176 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
										if v176 == int32(0) {
											v196 = v173
											*(*int32)(unsafe.Add(mBase, uint32(v196)+12)) = v145
										} else {
											if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v176))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v145)) == int32(0) {
												v190 = base.B2i32(base.Ui32(v145) < base.Ui32(v176))
											} else {
												v190 = int32(base.Ui32(v145-v176) >> (uint(int32(31)) % 32))
											}
											if v190 == int32(0) {
											} else {
												v194 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
												v196 = v194
												*(*int32)(unsafe.Add(mBase, uint32(v196)+12)) = v145
											}
										}
									}
								}
								v201 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
								F_LWLockRelease(m, v201+int32(_a_F_predicatelock_twophase_recover_5))
								mBase = m.M
								v205 = m.ExcPending
								if v205 != 0 {
									return
								} else {
									v217 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
									F_LWLockRelease(m, v217+int32(3584))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v140 = v121
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v140)+16)) = v141
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
						v147 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
						v151 = F_LWLockAcquire(m, v147+int32(_a_F_predicatelock_twophase_recover_5), int32(0))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return
						} else {
							if v145 == int32(0) {
								v156 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
								*(*int64)(unsafe.Add(mBase, uint32(v156)+8)) = int64(0)
							} else {
								v161 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])))
								if v161 == int32(1) {
									v166 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[5]))
									v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+316))
									v169 = base.B2i32(v167 != int32(2))
									*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])) = uint8(v169)
									v171 = v169
								} else {
									v171 = int32(0)
								}
								v173 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
								if v171 == int32(0) {
									v196 = v173
									*(*int32)(unsafe.Add(mBase, uint32(v196)+12)) = v145
								} else {
									v176 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
									if v176 == int32(0) {
										v196 = v173
										*(*int32)(unsafe.Add(mBase, uint32(v196)+12)) = v145
									} else {
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v176))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v145)) == int32(0) {
											v190 = base.B2i32(base.Ui32(v145) < base.Ui32(v176))
										} else {
											v190 = int32(base.Ui32(v145-v176) >> (uint(int32(31)) % 32))
										}
										if v190 == int32(0) {
										} else {
											v194 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
											v196 = v194
											*(*int32)(unsafe.Add(mBase, uint32(v196)+12)) = v145
										}
									}
								}
							}
							v201 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
							F_LWLockRelease(m, v201+int32(_a_F_predicatelock_twophase_recover_5))
							mBase = m.M
							v205 = m.ExcPending
							if v205 != 0 {
								return
							} else {
								v217 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
								F_LWLockRelease(m, v217+int32(3584))
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	case 1:
		v223 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[6]))
		v225 = l2 + int32(4)
		v226 = F_get_hash_value(m, v223, v225)
		mBase = m.M
		v227 = m.ExcPending
		if v227 != 0 {
			return
		} else {
			v229 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
			v233 = F_LWLockAcquire(m, v229+int32(3584), int32(1))
			mBase = m.M
			v234 = m.ExcPending
			if v234 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
				v237 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[2]))
				v240 = int32(0)
				v242 = F_hash_search(m, v237, v10+int32(4), v240, v240)
				mBase = m.M
				v243 = m.ExcPending
				if v243 != 0 {
					return
				} else {
					v245 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
					F_LWLockRelease(m, v245+int32(3584))
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return
					} else {
						v250 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
						F_CreatePredicateLock(m, v225, v226, v250)
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		}
	default:
		m.G0 = v10 + int32(16)
		return
	}
}
func F_prepare_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_prepare_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(993)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v15
	v19 = int32(_a_F_prepare_cb_wrapper_1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_prepare_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_prepare_cb_wrapper[0])) = v9 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+147)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+164)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if v37 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_prepare_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_prepare_cb_wrapper_3), v9)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_prepare_cb_wrapper_4), int32(986), int32(_a_F_prepare_cb_wrapper_5))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.T0[v37].(func(*base.Module, int32, int32, int64))(m, v11, l1, l2)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_prepare_cb_wrapper[0])) = v60
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_printsimple(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v136 int32
	_ = v136
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v15 < v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_slot_getsomeattrs_int(m, l0, v14)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v22 = v9 + int32(-16)
	F_pq_beginmessage(m, v22, int32(68))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	F_enlargeStringInfo(m, v22, int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v33 = int32(8)
	v37 = v26<<(uint(v33)%32) | int32(base.Ui32(v26)>>(uint(v33)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v31))) = uint16(v37)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v30 + int32(2)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if int32(0) < v42 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = v42
	v50 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_pq_endmessage(m, v9+int32(-16))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L70
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v50))))
	if v56 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L10
L13:
	;
	v319 = v50 + int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v319 < v320 {
		v49 = v320
		v50 = v319
		goto L11
	} else {
		goto L69
	}
L14:
	;
	F_enlargeStringInfo(m, v9+int32(-16), int32(4))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v50<<(uint(int32(2))%32))))
	v82 = v13 + v49<<(uint(int32(4))%32) + v50*int32(100)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+88))
	switch v83 - int32(20) {
	case 0:
		goto L21
	default:
		goto L19
	case 3:
		goto L18
	case 5:
		goto L22
	case 6:
		goto L20
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v64+v65))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v64 + int32(4)
	goto L13
L18:
	;
	v294 = v9 + int32(-48)
	if int32(0) <= v76 {
		goto L65
	} else {
		goto L66
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L61
	}
L20:
	;
	v154 = v9 + int32(-48)
	v155 = int32(0)
	if v76 == v155 {
		goto L45
	} else {
		goto L46
	}
L21:
	;
	v130 = v9 + int32(-48)
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
	if int64(0) <= v131 {
		goto L40
	} else {
		goto L41
	}
L22:
	;
	v88 = F_pg_detoast_datum_packed(m, v76)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v90 = int32(1)
	v91 = v88 + v90
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	v96 = v94 & v90
	if v96 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v97 = v91
	goto L26
L25:
	;
	v97 = v88 + int32(4)
	goto L26
L26:
	;
	if v94 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_pq_sendcountedtext(m, v9+int32(-16), v97, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L38
	}
L28:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v103 == int32(18) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v114 = int32(1)
	if v96 != 0 {
		v124 = int32(base.Ui32(v94)>>(uint(v114)%32)) - v114
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v106 = int32(16)
	goto L33
L32:
	;
	v106 = int32(0)
	goto L33
L33:
	;
	if base.Ui32((v103-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v113 = int32(4)
	goto L36
L35:
	;
	v113 = v106
	goto L36
L36:
	;
	v124 = v113
	goto L27
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v124 = int32(base.Ui32(v118)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	goto L13
L39:
	;
	F_pq_sendcountedtext(m, v9+int32(-16), v130, v145)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L43
	}
L40:
	;
	v141 = v131
	v142 = int32(0)
	goto L42
L41:
	;
	v136 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v136)
	v141 = int64(0) - v131
	v142 = int32(1)
	goto L42
L42:
	;
	v144 = F_pg_ulltoa_n(m, v141, v130+v142)
	mBase = m.M
	v145 = v144 + v142
	v147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v130+v145))) = uint8(v147)
	goto L39
L43:
	;
	goto L13
L44:
	;
	F_pq_sendcountedtext(m, v9+int32(-16), v154, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L60
	}
L45:
	;
	v164 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v164)
	v274 = int32(1)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v170 = int32(1233)
	v175 = int32(base.Ui32((base.I32_clz(v76)^int32(31))*v170+v170) >> (uint(int32(12)) % 32))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175<<(uint(int32(2))%32))+uint32(_c_F_printsimple[0])))
	v180 = v175 + base.B2i32(base.Ui32(v178) <= base.Ui32(v76))
	if base.Ui32(int32(_a_F_printsimple_0)) <= base.Ui32(v76) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v184 = v76
	v186 = v155
	goto L51
L49:
	;
	v220 = v76
	v222 = v155
	goto L50
L50:
	;
	if base.Ui32(int32(100)) <= base.Ui32(v220) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v193 = v154 + v180 - v186
	v194 = int32(4)
	v197 = base.I32_div_u_s(v184, int32(_a_F_printsimple_0))
	v200 = v184 + v197*int32(-10000)
	v201 = int32(100)
	v202 = base.I32_div_u_s(v200, v201)
	v203 = int32(1)
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202<<(uint(v203)%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v193-v194))) = uint16(v205)
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v200-v202*v201)<<(uint(v203)%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v193-int32(2)))) = uint16(v214)
	v217 = v186 + v194
	if base.Ui32(int32(99999999)) < base.Ui32(v184) {
		v184 = v197
		v186 = v217
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v220 = v197
	v222 = v217
	goto L50
L53:
	;
	goto L52
L54:
	;
	v233 = int32(2)
	v235 = int32(_a_F_printsimple_1)
	v237 = int32(100)
	v238 = base.I32_div_u_s(v220&v235, v237)
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v220-v238*v237)&v235<<(uint(int32(1))%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v154+v180-v222-v233))) = uint16(v246)
	v250 = v238
	v251 = v222 | v233
	goto L56
L55:
	;
	v250 = v220
	v251 = v222
	goto L56
L56:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v250) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v250<<(uint(int32(1))%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v154+v180-v251-int32(2)))) = uint16(v260)
	v274 = v180
	goto L44
L58:
	;
	goto L59
L59:
	;
	v263 = v250 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v263)
	v274 = v180
	goto L44
L60:
	;
	goto L13
L61:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v82)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v281
	F_errmsg_internal(m, int32(_a_F_printsimple_2), v11)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_printsimple_3), int32(136), int32(_a_F_printsimple_4))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_pq_sendcountedtext(m, v9+int32(-16), v294, v308)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L68
	}
L65:
	;
	v304 = v76
	v305 = int32(0)
	goto L67
L66:
	;
	v299 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v294))) = uint8(v299)
	v304 = int32(0) - v76
	v305 = int32(1)
	goto L67
L67:
	;
	v307 = F_pg_ultoa_n(m, v304, v294+v305)
	mBase = m.M
	v308 = v307 + v305
	v310 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v294+v308))) = uint8(v310)
	goto L64
L68:
	;
	goto L13
L69:
	;
	goto L12
L70:
	;
	m.G0 = v11 - int32(-64)
	return int32(1)
}
func F_privilege_to_string(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v1 = l0
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if v1&(v1-int64(1)) == int64(0) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(base.I64_ctz(v1))<<(uint(int32(2))%32))+uint32(_c_F_privilege_to_string[0])))
		m.G0 = v6 + int32(16)
		return v17
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*uint32)(unsafe.Add(mBase, uint32(v6))) = uint32(v1)
			F_errmsg_internal(m, int32(_a_F_privilege_to_string_0), v6)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_privilege_to_string_1), int32(2643), int32(_a_F_privilege_to_string_2))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_processPendingPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v5
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(-1)
	if base.Ui32(int32(25)) <= base.Ui32(v19) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = int32(base.Ui32(v19+int32(_a_F_processPendingPage_0)) >> (uint(int32(2)) % 32))
	goto L3
L2:
	;
	v31 = v5
	goto L3
L3:
	;
	v33 = v31 & int32(_a_F_processPendingPage_1)
	if base.Ui32(l3) <= base.Ui32(v33) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = l3
	v46 = v5
	goto L7
L5:
	;
	v151 = v5
	v154 = v5
	goto L6
L6:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_ginInsertBAEntries(m, l0, v15+int32(8), v154&int32(_a_F_processPendingPage_1), v161, v162, v151)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L33
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(20)+v40&int32(_a_F_processPendingPage_1)<<(uint(int32(2))%32))))
	v58 = l2 + v55&int32(_a_F_processPendingPage_2)
	v59 = F_gintuple_get_attrnum(m, v49, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v151 = v138
	v154 = v101
	goto L6
L9:
	;
	return
L10:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)))
	if v61 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v105 = F_gintuple_get_key(m, v102, v58, v15+int32(7))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L26
	}
L12:
	;
	v63 = v15 + int32(8)
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+2)))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63))))
	v66 = int32(16)
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+2)))
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58))))
	if v64|v65<<(uint(v66)%32) == v69|v70<<(uint(v66)%32) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	goto L14
L14:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v95)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v97
	v101 = v59
	goto L11
L15:
	;
	v83 = v46 & int32(_a_F_processPendingPage_1)
	if v59 == v83 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	goto L15
L17:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)))
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
	if v76 == v77 {
		v80 = int32(1)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v80 = int32(0)
	goto L16
L20:
	;
	goto L19
L21:
	;
	v85 = v80
	goto L23
L22:
	;
	v85 = int32(0)
	goto L23
L23:
	;
	if v85 != 0 {
		v101 = v46
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_ginInsertBAEntries(m, l0, v63, v83, v86, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	goto L14
L26:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v109 <= v108 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v109 << (uint(int32(1)) % 32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v117 = F_repalloc(m, v114, v109<<(uint(int32(3))%32))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	v126 = v108
	goto L29
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v127+v126<<(uint(int32(2))%32)))) = v105
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v132+v133))) = uint8(v107)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v137 = int32(1)
	v138 = v136 + v137
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v138
	v141 = v40 + v137
	if base.Ui32(v141&int32(_a_F_processPendingPage_1)) <= base.Ui32(v33) {
		v40 = v141
		v46 = v101
		goto L7
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v117
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v122 = F_repalloc(m, v120, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v126 = v125
	goto L29
L32:
	;
	goto L8
L33:
	;
	m.G0 = v15 + int32(16)
	return
}
func F_process_syncing_tables(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int64
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v329 int32
	_ = v329
	var v332 int64
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int64
	_ = v604
	var v605 int64
	_ = v605
	var v613 int64
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int64
	_ = v624
	var v626 int32
	_ = v626
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(160)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	switch v22 {
	case 0:
		goto L3
	case 1:
		goto L5
	case 2:
		goto L4
	default:
		goto L1
	}
L1:
	;
	m.G0 = v18 + int32(160)
	return
L2:
	;
	v763 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v36)+56)), uint32(v763))
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L9
	} else {
		goto L181
	}
L4:
	;
	v154 = F_FetchTableStates(m, v18+int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L29
	}
L5:
	;
	v25 = base.AtomicRmwXchg32(m, v21, int32(56), int32(1))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	F_s_lock(m, v27+int32(56), int32(_a_F_process_syncing_tables_0), int32(296), int32(_a_F_process_syncing_tables_1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+40)))
	if v37 != int32(99) {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v36)+48))
	if base.Ui64(l0) < base.Ui64(v40) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v42 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+152)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = l0
	v75 = int32(115)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+40)) = uint8(v75)
	v77 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v36)+56)), uint32(v77))
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[1]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	goto L13
L13:
	;
	if base.B2i32(v82 == int32(2)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+32))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+36))
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v90)+40)))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v90)+48))
	F_UpdateSubscriptionRelState(m, v91, v92, v93, v94, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[2]))
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[3]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+36))
	m.T0[v104].(func(*base.Module, int32, int32))(m, v99, v18+int32(92))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+32))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+36))
	v112 = v18 + int32(96)
	F_ReplicationSlotNameForTablesync(m, v109, v110, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[2]))
	F_ReplicationSlotDropAtPubNode(m, v116, v112, int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v123 = F_pgstat_report_stat(m, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+32))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v128)+36))
	v132 = v18 + int32(16)
	F_ReplicationOriginNameForLogicalRep(m, v129, v130, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_replorigin_session_reset(m)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v138 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_process_syncing_tables[4])) = v138
	*(*int64)(unsafe.Add(mBase, _c_F_process_syncing_tables[5])) = v138
	v144 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _c_F_process_syncing_tables[6])) = uint16(v144)
	F_replorigin_drop_by_name(m, v132, int32(1), v144)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	F_finish_sync_worker(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[7]))
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[8]))
	v160 = int32(0)
	if v157|base.B2i32(v159 == v160) == v160 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[8]))
	if v185 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = int64(68719476740)
	v173 = F_hash_create(m, int32(_a_F_process_syncing_tables_2), int32(256), v18+int32(96), int32(40))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v159|base.B2i32(v157 == int32(0)) != 0 {
		goto L30
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[7])) = v173
	goto L30
L35:
	;
	F_hash_destroy(m, v157)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[7])) = int32(0)
	goto L30
L37:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	if v687 != int32(1) {
		goto L1
	} else {
		goto L158
	}
L38:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v188 <= int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v196 = v2
	v200 = v2
	goto L40
L40:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206+v200<<(uint(int32(2))%32))))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+16)))
	if v211 == int32(115) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v653 == int32(0) {
		goto L37
	} else {
		goto L156
	}
L42:
	;
	v664 = v200 + int32(1)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if v664 < v665 {
		v196 = v653
		v200 = v664
		goto L40
	} else {
		goto L155
	}
L43:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v210)+8))
	if base.Ui64(l0) < base.Ui64(v214) {
		v653 = v196
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	v266 = F_LWLockAcquire(m, v262+int32(_a_F_process_syncing_tables_3), int32(1))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L9
	} else {
		goto L59
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v210)+8)) = l0
	v217 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v210)+16)) = uint8(v217)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	if v219 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L9
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+32))
	F_LockSharedObject(m, int32(_a_F_process_syncing_tables_4), v229, int32(1))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L9
	} else {
		goto L51
	}
L50:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)) = uint8(v224)
	goto L49
L51:
	;
	if v196 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v237 = F_table_open(m, int32(_a_F_process_syncing_tables_5), int32(3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L9
	} else {
		goto L55
	}
L53:
	;
	v239 = v196
	goto L54
L54:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+32))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v245 = v18 + int32(96)
	F_ReplicationOriginNameForLogicalRep(m, v242, v243, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L9
	} else {
		goto L56
	}
L55:
	;
	v239 = v237
	goto L54
L56:
	;
	F_replorigin_drop_by_name(m, v245, int32(1), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+32))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v256 = int32(*(*int8)(unsafe.Add(mBase, uint32(v210)+16)))
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v210)+8))
	F_UpdateSubscriptionRelState(m, v254, v255, v256, v257, int32(1))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v653 = v239
	goto L42
L59:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+32))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v272 = int32(0)
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[10]))
	if v277 <= v272 {
		v309 = v272
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v309 != 0 {
		goto L71
	} else {
		goto L72
	}
L61:
	;
	goto L60
L62:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[11]))
	v287 = v272
	goto L63
L63:
	;
	v292 = v281 + int32(16) + v287*int32(112)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+16)))
	if v293 != int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v309 = int32(0)
	goto L61
L65:
	;
	v304 = v287 + int32(1)
	if v304 != v277 {
		v287 = v304
		goto L63
	} else {
		goto L70
	}
L66:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	if v296 == int32(3) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v292)+32))
	if v299 != v270 {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v292)+36))
	if v301 != v271 {
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v309 = v292
	goto L61
L70:
	;
	goto L64
L71:
	;
	v313 = int32(56)
	v314 = v309 + v313
	v317 = base.AtomicRmwXchg32(m, v309, v313, int32(1))
	if v317 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+32))
	v483 = int32(0)
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[10]))
	if v487 <= v483 {
		v575 = v483
		goto L128
	} else {
		goto L129
	}
L74:
	;
	F_s_lock(m, v314, int32(_a_F_process_syncing_tables_0), int32(537), int32(_a_F_process_syncing_tables_6))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L9
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v210)+16)) = uint8(v323)
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v309)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v210)+8)) = v325
	if v323 == int32(119) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v329 = int32(99)
	*(*uint8)(unsafe.Add(mBase, uint32(v309)+40)) = uint8(v329)
	if base.Ui64(l0) < base.Ui64(v325) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v334 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v314))), uint32(v334))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+16)))
	if v337 == int32(119) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v332 = v325
	goto L83
L82:
	;
	v332 = l0
	goto L83
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v309)+48)) = v332
	goto L80
L84:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	if v340 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	F_LWLockRelease(m, v475+int32(_a_F_process_syncing_tables_3))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L9
	} else {
		goto L127
	}
L87:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	F_SetLatch(m, v341+int32(20))
	mBase = m.M
	goto L90
L88:
	;
	goto L89
L89:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	F_LWLockRelease(m, v346+int32(_a_F_process_syncing_tables_3))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L9
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	if v351 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v196 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v363 = v196
	goto L94
L94:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L9
	} else {
		goto L101
	}
L95:
	;
	F_relation_close(m, v196, int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L9
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v359 = int32(0)
	v361 = F_pgstat_report_stat(m, v359)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	v363 = v359
	goto L94
L101:
	;
	v366 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)) = uint8(v366)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	goto L102
L102:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[12]))
	if v385 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L9
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L9
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+32))
	v395 = F_GetSubscriptionRelState(m, v392, v368, v18+int32(96))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	if base.B2i32(v395 == int32(0))|base.B2i32(v395&int32(255) == int32(115)) != 0 {
		v653 = v363
		goto L42
	} else {
		goto L110
	}
L110:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	v409 = F_LWLockAcquire(m, v405+int32(_a_F_process_syncing_tables_3), int32(1))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+32))
	v414 = int32(0)
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[10]))
	if v419 <= v414 {
		v451 = v414
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	F_LWLockRelease(m, v456+int32(_a_F_process_syncing_tables_3))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L9
	} else {
		goto L123
	}
L113:
	;
	goto L112
L114:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[11]))
	v429 = v414
	goto L115
L115:
	;
	v434 = v423 + int32(16) + v429*int32(112)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+16)))
	if v435 != int32(1) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v451 = int32(0)
	goto L113
L117:
	;
	v446 = v429 + int32(1)
	if v446 != v419 {
		v429 = v446
		goto L115
	} else {
		goto L122
	}
L118:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	if v438 == int32(3) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v434)+32))
	if v441 != v413 {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v434)+36))
	if v443 != v368 {
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v451 = v434
	goto L113
L122:
	;
	goto L116
L123:
	;
	if v451 == int32(0) {
		v653 = v363
		goto L42
	} else {
		goto L124
	}
L124:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[13]))
	v468 = F_WaitLatch(m, v464, int32(41), int32(1000), int32(134217760))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = int32(0)
	goto L126
L126:
	;
	goto L102
L127:
	;
	v653 = v196
	goto L42
L128:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[9]))
	F_LWLockRelease(m, v588+int32(_a_F_process_syncing_tables_3))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L9
	} else {
		goto L145
	}
L129:
	;
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[11]))
	v493 = v491 + int32(16)
	if v487 != int32(1) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v503 = v483
	v506 = v483
	v510 = v483
	goto L133
L131:
	;
	v548 = v483
	v551 = v483
	goto L132
L132:
	;
	v562 = v493 + v551*int32(112)
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+16)))
	if v563 != int32(1) {
		v575 = v548
		goto L128
	} else {
		goto L143
	}
L133:
	;
	v517 = v493 + v506*int32(112)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+16)))
	if v518 != int32(1) {
		v527 = v503
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v487&int32(1) == int32(0) {
		v575 = v537
		goto L128
	} else {
		goto L142
	}
L135:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+128)))
	if v528 != int32(1) {
		v537 = v527
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	if v521 != int32(1) {
		v527 = v503
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v517)+32))
	v527 = v503 + base.B2i32(v524 == v482)
	goto L135
L138:
	;
	v538 = int32(2)
	v539 = v506 + v538
	v541 = v510 + v538
	if v541 != v487&int32(2147483646) {
		v503 = v537
		v506 = v539
		v510 = v541
		goto L133
	} else {
		goto L141
	}
L139:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v517)+112))
	if v531 != int32(1) {
		v537 = v527
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v517)+144))
	v537 = v527 + base.B2i32(v534 == v482)
	goto L138
L141:
	;
	goto L134
L142:
	;
	v548 = v537
	v551 = v539
	goto L132
L143:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	if v566 != int32(1) {
		v575 = v548
		goto L128
	} else {
		goto L144
	}
L144:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v562)+32))
	v575 = v548 + base.B2i32(v569 == v482)
	goto L128
L145:
	;
	v594 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[14]))
	if v594 <= v575 {
		v653 = v196
		goto L42
	} else {
		goto L146
	}
L146:
	;
	v599 = m.G0
	v600 = int32(16)
	v601 = v599 - v600
	m.G0 = v601
	F_gettimeofday(m, v601)
	mBase = m.M
	v604 = *(*int64)(unsafe.Add(mBase, uint32(v601)))
	v605 = int64(*(*int32)(unsafe.Add(mBase, uint32(v601)+8)))
	m.G0 = v601 + v600
	v613 = v605 + v604*int64(1000000) - int64(946684800000000)
	goto L147
L147:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[7]))
	v619 = F_hash_search(m, v615, v210, int32(1), v18+int32(96))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L9
	} else {
		goto L148
	}
L148:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+96)))
	if v621 == int32(1) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v624 = *(*int64)(unsafe.Add(mBase, uint32(v619)+8))
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[15]))
	goto L152
L150:
	;
	goto L151
L151:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v619)+8)) = v613
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[0]))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+24))
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[16]))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v640)+16))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v637)+28))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v646 = F_logicalrep_worker_launch(m, int32(1), v638, v641, v642, v643, v644, int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L9
	} else {
		goto L154
	}
L152:
	;
	if base.B2i32(base.I64_extend_i32_s(v626)*int64(1000) <= v613-v624) == int32(0) {
		v653 = v196
		goto L42
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v653 = v196
	goto L42
L155:
	;
	goto L41
L156:
	;
	F_relation_close(m, v653, int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	goto L37
L158:
	;
	v691 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[16]))
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691)+28)))
	if v692 != int32(112) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L9
	} else {
		goto L179
	}
L160:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L9
	} else {
		goto L161
	}
L161:
	;
	v699 = F_FetchTableStates(m, v18+int32(96))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L9
	} else {
		goto L162
	}
L162:
	;
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+96)))
	if v701 == int32(1) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L9
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[8]))
	v711 = int32(0)
	if v699&base.B2i32(v710 == v711) == v711 {
		goto L159
	} else {
		goto L168
	}
L166:
	;
	v707 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v718 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L9
	} else {
		goto L169
	}
L169:
	;
	if v718 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[16]))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v721)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v722
	F_errmsg(m, int32(_a_F_process_syncing_tables_7), v18)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L9
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L9
	} else {
		goto L175
	}
L173:
	;
	F_errfinish(m, int32(_a_F_process_syncing_tables_0), int32(670), int32(_a_F_process_syncing_tables_6))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L9
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v735 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L9
	} else {
		goto L176
	}
L176:
	;
	v738 = *(*int32)(unsafe.Add(mBase, _c_F_process_syncing_tables[16]))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v738)))
	F_ApplyLauncherForgetWorkerStartTime(m, v739)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L9
	} else {
		goto L177
	}
L177:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	v748 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	goto L1
L181:
	;
	F_errmsg_internal(m, int32(_a_F_process_syncing_tables_8), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L9
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_process_syncing_tables_0), int32(718), int32(_a_F_process_syncing_tables_9))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L9
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_prsd_lextype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v6 = F_palloc(m, int32(288))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = int32(1)
	goto L3
L3:
	;
	v15 = int32(12)
	v17 = v6 + v11*v15
	*(*int32)(unsafe.Add(mBase, uint32(v17-v15))) = v11
	v24 = v11 << (uint(int32(2)) % 32)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_prsd_lextype[0])))
	v26 = F_pstrdup(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+276)) = int32(0)
	return v6
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17-int32(8)))) = v26
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_prsd_lextype[1])))
	v32 = F_pstrdup(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17-int32(4)))) = v32
	v36 = v11 + int32(1)
	if v36 != int32(24) {
		v11 = v36
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
}
func F_pt_contained_poly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v11 = F_point_inside(m, v2, v8, v4+int32(40))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v11 != int32(0))
		}
	}
}
func F_pull_up_union_leaf_queries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = l0
	goto L1
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v25 != int32(142) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L29
	}
L3:
	;
	goto L2
L4:
	;
	if v25 != int32(63) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_pull_up_union_leaf_queries(m, v123, l1, l2, l3, l4)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L28
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v32 = F_palloc0(m, int32(36))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+12)) = int64(0)
	v36 = v30 + l4
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(322)
	v41 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v45 = v44
	goto L12
L11:
	;
	v45 = v41
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v45
	v49 = F_palloc0(m, v45<<(uint(int32(1))%32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v52 == int32(0) {
		v98 = v41
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v98
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v108 = F_lappend(m, v107, v32)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L25
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v55 <= int32(0) {
		v98 = v41
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v59 = int32(0)
	v61 = v55
	v63 = v41
	goto L17
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v59<<(uint(int32(2))%32))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+26)))
	if v74 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v98 = v90
	goto L14
L19:
	;
	v77 = F_makeVarFromTargetEntry(m, v36, v73)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	v89 = v61
	v90 = v63
	goto L21
L21:
	;
	v92 = v59 + int32(1)
	if v92 < v89 {
		v59 = v92
		v61 = v89
		v63 = v90
		goto L17
	} else {
		goto L24
	}
L22:
	;
	v79 = F_lappend(m, v63, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v49+v81<<(uint(int32(1))%32)-int32(2)))) = uint16(v81)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v89 = v88
	v90 = v79
	goto L21
L24:
	;
	goto L18
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v108
	v112 = F_palloc0(m, int32(8))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = int32(63)
	v118 = F_pull_up_subqueries_recurse(m, l1, v112, int32(0), v32)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v13 + int32(16)
	return
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v15 = v126
	goto L1
L29:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v131
	F_errmsg_internal(m, int32(_a_F_pull_up_union_leaf_queries_0), v13)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_pull_up_union_leaf_queries_1), int32(1756), int32(_a_F_pull_up_union_leaf_queries_2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pull_varattnos_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = int32(0)
	if l0 == v3 {
		v27 = v3
		return v27
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 == int32(6) {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v9 != v10 {
				v27 = v3
				return v27
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v12 != 0 {
					v27 = v3
					return v27
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
					v17 = F_bms_add_member(m, v13, v14+int32(7))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
						return int32(0)
					}
				}
			}
		} else {
			v25 = F_expression_tree_walker_impl(m, l0, int32(897), l1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = v25
				return v27
			}
		}
	}
}
func F_pullf_create_mbuf_reader(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14001(m, l0, l1, int32(_a_F_pullf_create_mbuf_reader_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pullf_read_fixed(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pullf_read_max(m, l0, l1, v8+int32(12), l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 < int32(0) {
			v32 = v12
			m.G0 = v8 + int32(16)
			return v32
		} else {
			if v12 != l1 {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
				F_px_debug(m, int32(_a_F_pullf_read_fixed_0), v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v32 = int32(-100)
					m.G0 = v8 + int32(16)
					return v32
				}
			} else {
				v25 = int32(0)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				if base.B2i32(l1 == v25)|base.B2i32(v28 == l2) != 0 {
					v32 = v25
				} else {
					base.MemoryCopy(m, l2, v28, l1)
					v32 = v25
				}
				m.G0 = v8 + int32(16)
				return v32
			}
		}
	}
}
func F_pushValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	v4 = l3
	v5 = l4
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l2 <= int32(2046) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return
L2:
	;
	v174 = l2 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v177 = v175 - v176
	v178 = v174 + v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v179 <= v178 {
		goto L32
	} else {
		goto L33
	}
L3:
	;
	v154 = F_palloc0(m, int32(12))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L19
	} else {
		goto L30
	}
L4:
	;
	if l2 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v136 = F_errsave_start(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L19
	} else {
		goto L25
	}
L7:
	;
	v18 = int32(-1)
	if l2 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v109 = int32(0)
	goto L9
L9:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v112 = v110 - v111
	if v112 < int32(_a_F_pushValue_0) {
		goto L3
	} else {
		goto L18
	}
L10:
	;
	v109 = v89 ^ int32(-1)
	goto L9
L11:
	;
	v30 = v18
	v31 = l1
	v32 = int32(0)
	goto L14
L12:
	;
	v68 = v18
	v69 = l1
	goto L13
L13:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32((v74^int32(base.Ui32(v68)>>(uint(int32(24))%32)))<<(uint(int32(2))%32))+uint32(_c_F_pushValue[0])))
	v89 = v80 ^ v68<<(uint(int32(8))%32)
	goto L10
L14:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v38 = int32(24)
	v41 = int32(2)
	v43 = *(*int32)(unsafe.Add(mBase, uint32((v37^int32(base.Ui32(v30)>>(uint(v38)%32)))<<(uint(v41)%32))+uint32(_c_F_pushValue[0])))
	v44 = int32(8)
	v46 = v43 ^ v30<<(uint(v44)%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32((v36^int32(base.Ui32(v46)>>(uint(v38)%32)))<<(uint(v41)%32))+uint32(_c_F_pushValue[0])))
	v55 = v52 ^ v46<<(uint(v44)%32)
	v57 = v31 + v41
	v59 = v32 + v41
	if v59 != l2&int32(-2) {
		v30 = v55
		v31 = v57
		v32 = v59
		goto L14
	} else {
		goto L16
	}
L15:
	;
	if l2&int32(1) == int32(0) {
		v89 = v55
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v68 = v55
	v69 = v57
	goto L13
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v116 = F_errsave_start(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	if v116 == int32(0) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v123
	F_errmsg(m, int32(_a_F_pushValue_1), v14+int32(16))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	F_errsave_finish(m, v115, int32(_a_F_pushValue_2), int32(555), int32(_a_F_pushValue_3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L2
L25:
	;
	if v136 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v143
	F_errmsg(m, int32(_a_F_pushValue_4), v14)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, v135, int32(_a_F_pushValue_2), int32(588), int32(_a_F_pushValue_5))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	goto L1
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v109
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)) = uint8(v4)
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = l2&int32(4095) | v112<<(uint(int32(12))%32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v168 = F_lcons(m, v154, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v168
	goto L2
L32:
	;
	v186 = v179
	v187 = v176
	goto L35
L33:
	;
	v209 = v175
	goto L34
L34:
	;
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v193 = v186 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v193
	v195 = F_repalloc(m, v187, v193)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L19
	} else {
		goto L37
	}
L36:
	;
	v209 = v198
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v195
	v198 = v177 + v195
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v200 <= v178 {
		v186 = v200
		v187 = v195
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	base.MemoryCopy(m, v209, l1, l2)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v215 = v214 + l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v215
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v217)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v219 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v174 + v223
	goto L1
}
func F_pushf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 != 0 {
		v14 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l3, l2, v9+int32(12))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if int32(0) <= v14 {
				v22 = v14
				v24 = F_palloc0(m, int32(24))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
					if v22 != 0 {
						v32 = F_palloc(m, v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = v32
							v35 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
							v41 = v35
							m.G0 = v9 + int32(16)
							return v41
						}
					} else {
						v34 = int32(0)
						v35 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
						v41 = v35
						m.G0 = v9 + int32(16)
						return v41
					}
				}
			} else {
				v41 = v14
				m.G0 = v9 + int32(16)
				return v41
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
		v22 = int32(0)
		v24 = F_palloc0(m, int32(24))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
			if v22 != 0 {
				v32 = F_palloc(m, v22)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v32
					v35 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
					v41 = v35
					m.G0 = v9 + int32(16)
					return v41
				}
			} else {
				v34 = int32(0)
				v35 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
				v41 = v35
				m.G0 = v9 + int32(16)
				return v41
			}
		}
	}
}
func F_pushf_flush(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v54
L2:
	;
	v6 = l0
	goto L5
L3:
	;
	goto L4
L4:
	;
	v54 = int32(0)
	goto L1
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if int32(0) < v11 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v19 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v38 != 0 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	if int32(0) < v27 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v21 = m.T0[v19].(func(*base.Module, int32, int32, int32, int32) int32)(m, v14, v20, v16, v15)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v25 = F_pushf_write(m, v14, v16, v15)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	v27 = v21
	goto L10
L16:
	;
	v27 = v25
	goto L10
L17:
	;
	v30 = int32(-12)
	goto L19
L18:
	;
	v30 = v27
	goto L19
L19:
	;
	if v30 < int32(0) {
		v54 = v30
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v41 = m.T0[v38].(func(*base.Module, int32, int32) int32)(m, v39, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v46 != 0 {
		v6 = v46
		goto L5
	} else {
		goto L26
	}
L24:
	;
	if v41 < int32(0) {
		v54 = v41
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L6
}
