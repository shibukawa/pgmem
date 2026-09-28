package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintBufferIO(m *base.Module, l0 int64) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14221(m, l0, int32(_a_F_ResOwnerPrintBufferIO_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_ResOwnerPrintCatCacheList(m *base.Module, l0 int64) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = base.I32_wrap_i64(l0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
	v20 = F_psprintf(m, int32(_a_F_ResOwnerPrintCatCacheList_0), v8)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v20
	}
}
func F_ResOwnerReleaseBuffer(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = base.I32_wrap_i64(l0)
	if v9 != 0 {
		if v9 < int32(0) {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[0]))
			v15 = v9 ^ int32(-1)
			v18 = v13 + v15<<(uint(int32(2))%32)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v21 = v19 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v21
			if v21 == int32(0) {
				v25 = int32(_a_F_ResOwnerReleaseBuffer_0)
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[1])) = v27 - int32(1)
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[2]))
				v35 = v32 + v15*int32(56)
				v36 = int64(0)
				v39 = base.AtomicRmwCmpxchg64(m, v35, int32(24), v36, v36)
				*(*int64)(unsafe.Add(mBase, uint32(v35)+24)) = v39 - int64(1)
			} else {
			}
			m.G0 = v7 + int32(16)
			return
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[3]))
			if v45 != int32(-1) {
				v49 = v45 << (uint(int32(4)) % 32)
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_c_F_ResOwnerReleaseBuffer[4])))
				if v52 == v9 {
					v58 = v49 + int32(_a_F_ResOwnerReleaseBuffer_1)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
					if v59 != 0 {
						v60 = int32(_a_F_ResOwnerReleaseBuffer_2)
						v62 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[5]))
						*(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[5])) = v62 + int32(1)
						v67 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[6]))
						v68 = int32(56)
						F_BufferLockUnlock(m, v9, v67+v9*v68-v68)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[6]))
							v77 = int32(56)
							F_UnpinBufferNoOwner(m, v76+v9*v77-v77)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[6]))
						v77 = int32(56)
						F_UnpinBufferNoOwner(m, v76+v9*v77-v77)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				} else {
					v56 = F_GetPrivateRefCountEntrySlow(m, v9, int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						v58 = v56
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
						if v59 != 0 {
							v60 = int32(_a_F_ResOwnerReleaseBuffer_2)
							v62 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[5]))
							*(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[5])) = v62 + int32(1)
							v67 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[6]))
							v68 = int32(56)
							F_BufferLockUnlock(m, v9, v67+v9*v68-v68)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[6]))
								v77 = int32(56)
								F_UnpinBufferNoOwner(m, v76+v9*v77-v77)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[6]))
							v77 = int32(56)
							F_UnpinBufferNoOwner(m, v76+v9*v77-v77)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v56 = F_GetPrivateRefCountEntrySlow(m, v9, int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = v56
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
					if v59 != 0 {
						v60 = int32(_a_F_ResOwnerReleaseBuffer_2)
						v62 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[5]))
						*(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[5])) = v62 + int32(1)
						v67 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[6]))
						v68 = int32(56)
						F_BufferLockUnlock(m, v9, v67+v9*v68-v68)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[6]))
							v77 = int32(56)
							F_UnpinBufferNoOwner(m, v76+v9*v77-v77)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBuffer[6]))
						v77 = int32(56)
						F_UnpinBufferNoOwner(m, v76+v9*v77-v77)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v92 = m.ExcPending
		if v92 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
			F_errmsg_internal(m, int32(_a_F_ResOwnerReleaseBuffer_3), v7)
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ResOwnerReleaseBuffer_4), int32(_a_F_ResOwnerReleaseBuffer_5), int32(_a_F_ResOwnerReleaseBuffer_6))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_ResOwnerReleasePGMEMDigest(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = base.I32_wrap_i64(l0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	m.Env.Pgmem_hash_free(m, v7)
	mBase = m.M
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v9 != 0 {
		F_ResourceOwnerForget(m, v9, l0&int64(4294967295), int32(_a_F_ResOwnerReleasePGMEMDigest_0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_pfree(m, v4)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, v4)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F__equalResTarget(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v53
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = F_equal(m, v39, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	if v6 == int32(0) {
		v53 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 == v7 {
		goto L2
	} else {
		goto L15
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v34 != 0 {
		v53 = v3
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v18 = v7
	v19 = v6
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v33 = v23
	v34 = v22
	goto L8
L12:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L2
L15:
	;
	return int32(0)
L16:
	;
	return int32(0)
L17:
	;
	if v41 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	goto L20
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v51 = F_equal(m, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v53 = v51
	goto L1
}
